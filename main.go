package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
	"ti/base"
	_ "ti/builtin"
	"ti/context"
	"ti/eval"
	"ti/lexer"
	"ti/lexer/reader"
	"ti/lsp"
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

	if len(p.DefineInfos) > 0 && p.IsDefineInfo {
		for _, info := range p.DefineInfos {
			fmt.Println(info)
		}
	}

	if len(base.TSignatures) > 0 && p.IsDictOut {
		for _, sig := range base.TSignatures {
			fmt.Println("%" + sig.Contents + ":::" + sig.Detail)
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

// AnalyzeContent は文字列のRubyコードを解析します（LSPから使用）
func AnalyzeContent(content string, filename string) error {
	br := bufio.NewReader(strings.NewReader(content))

	for _, round := range context.GetRounds() {
		p := getParser(br, filename)
		cleanSimpleIdentifires()

		ctx := context.NewContext("", "", round)
		evaluator := eval.Evaluator{}
		p.Errors = []error{}

		for {
			t, err := p.Read()
			if err != nil {
				return err
			}

			err = evaluator.Eval(&p, ctx, t)
			if err != nil {
				return err
			}

			if t == nil {
				break
			}
		}

		// 次のラウンドのために br をリセット
		br = bufio.NewReader(strings.NewReader(content))
	}

	return nil
}

func main() {
	//	defer profile.Start().Stop()

	// Check for LSP mode
	if len(os.Args) >= 2 && os.Args[1] == "--lsp" {
		server := lsp.NewServer()
		server.RunStdio()
		return
	}

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

			if len(os.Args) > 0 && slices.Contains(os.Args, "-d") {
				p.Debug = true
			}

			if len(os.Args) > 0 && slices.Contains(os.Args, "-i") {
				p.IsDefineInfo = true
			}

			if len(os.Args) > 0 && slices.Contains(os.Args, "-a") {
				p.IsDictOut = true
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
