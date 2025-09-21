package main

import (
	"bufio"
	"fmt"
	"os"
	"ti/base"
	_ "ti/builtin"
	"ti/context"
	"ti/eval"
	"ti/lexer"
	"ti/lexer/reader"
	"ti/parser"
	"time"
	// "github.com/pkg/profile"
)

func getParser(br *bufio.Reader, file string) parser.Parser {
	lr := reader.New(*br)
	l := lexer.New(lr)

	return parser.New(l, file)
}

func loop(p parser.Parser, round string) {
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

	if len(p.DefineInfos) > 0 {
		for _, info := range p.DefineInfos {
			fmt.Println(info)
		}
	}

	if len(p.Errors) > 0 {
		for _, err := range p.Errors {
			fmt.Println(err)
		}
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

		for _, round := range context.GetRounds() {
			if len(os.Args) == 1 {
				panic("want one argument!")
			}

			file = os.Args[1]
			fp, _ := os.Open(file)
			br = bufio.NewReader(fp)

			p := getParser(br, file)

			if len(os.Args) >= 3 && os.Args[2] == "-d" {
				p.Debug = true
			}

			cleanSimpleIdentifires()

			loop(p, round)
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
