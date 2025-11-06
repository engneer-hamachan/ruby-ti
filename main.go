package main

import (
	"bufio"
	"fmt"
	"os"
	"ti/base"
	_ "ti/builtin"
	"ti/cmd"
	"ti/context"
	"ti/eval"
	"ti/lexer"
	"ti/lexer/reader"
	"ti/parser"
	"time"
)

func getParser(br *bufio.Reader, file string) parser.Parser {
	lr := reader.New(*br)
	l := lexer.New(lr)

	return parser.New(l, file)
}

func loop(p parser.Parser, flags *cmd.ExecuteFlags, round string) {
	ctx := context.NewContext("", "", round)
	evaluator := eval.Evaluator{}

	p.Errors = []error{}

	for {
		t, err := p.Read()
		if err != nil {
			p.Fatal(ctx, err)
		}

		err = evaluator.Eval(&p, ctx, t)
		if err != nil {
			p.Fatal(ctx, err)
		}

		if t != nil {
			continue
		}

		break
	}

	if round != "check" {
		return
	}

	if len(p.DefineInfos) > 0 && flags.IsDefineInfo {
		cmd.PrintDefineInfosForPlugin(p.DefineInfos)
	}

	if len(base.TSignatures) > 0 && flags.IsDefineAllInfo {
		cmd.PrintAllDefinitionsForLsp(p)
	}

	if len(base.TSignatures) > 0 && flags.IsLsp {
		cmd.PrintSuggestionsForLsp(p)
	}

	if len(p.Errors) > 0 {
		cmd.PrintAllErrorsForPlugin(p)
		os.Exit(0)
	}
}

func cleanSimpleIdentifires() {
	for key, value := range base.TFrame {
		if value.IsIdentifierType() && key.Variable() == value.ToString() {
			delete(base.TFrame, key)
		}
	}
}

func main() {
	//	defer profile.Start().Stop()

	timeout := time.After(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		var br *bufio.Reader
		var file string

		cmd.ValidateArgs()
		flags := cmd.BuildFlags()

		if flags.IsAllType {
			cmd.PrintAllTypes()
			done <- true
			return
		}

		for _, round := range context.GetRounds() {
			file = cmd.GetTargetFile()
			fp, _ := os.Open(file)
			br = bufio.NewReader(fp)

			p := getParser(br, file)
			cmd.ApplyParserFlags(&p)

			cleanSimpleIdentifires()

			loop(p, flags, round)
		}

		if flags.IsExtends {
			className := cmd.GetTargetClassName()
			cmd.PrintExtends(className)
		}

		done <- true
	}()

	select {
	case <-done:
		return

	case <-timeout:
		fmt.Println("timeout")
		os.Exit(1)
	}
}
