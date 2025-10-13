package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"ti/base"
	_ "ti/builtin"
	"ti/context"
	"ti/eval"
	"ti/lexer"
	"ti/lexer/reader"
	"ti/lsp"
	"ti/parser"
	"time"
	"unicode"
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

	if len(p.DefineInfos) > 0 && p.IsDefineInfo && round == "check" {
		for _, info := range p.DefineInfos {
			fmt.Println(info)
		}
	}

	if len(base.TSignatures) > 0 && p.IsLsp && round == "check" {
		frame := p.LspSudjestTargetT.GetFrame()
		if frame == "" {
			frame = "unknown"
		}
		fmt.Println("@" + frame + ":::" + p.LspSudjestTargetT.GetObjectClass())

		for _, sig := range base.TSignatures {
			objectClass := p.LspSudjestTargetT.GetObjectClass()
			if objectClass == "Identifier" {
				objectClass = ""
			}

			if objectClass == "" && slices.Contains([]string{"", "Kernel"}, sig.Class) {
				fmt.Println("%" + sig.Contents + ":::" + sig.Detail)
				continue
			}

			if isSudjest(objectClass, sig) {
				tmp := p.LspSudjestTargetT.GetBeforeEvaluateCode()
				if len(tmp) > 0 && unicode.IsUpper(rune(tmp[0])) == sig.IsStatic {
					fmt.Println("%" + sig.Contents + ":::" + sig.Detail)
				}
			}
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

func isSudjest(objectClass string, sig base.Sig) bool {
	if sig.Class == "" {
		return false
	}

	if sig.Class == objectClass {
		return true
	}

	return isParentClass(sig.Frame, objectClass, sig.Class)
}

func isParentClass(frame, childClass, parentClass string) bool {
	classNode := base.ClassNode{Frame: frame, Class: childClass}

	for _, parentNode := range base.ClassInheritanceMap[classNode] {
		if parentNode.Class == parentClass {
			return true
		}

		// Recursively check parent's parents
		if isParentClass(parentNode.Frame, parentNode.Class, parentClass) {
			return true
		}
	}

	return false
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
				p.IsLsp = true
				if len(os.Args) > 3 {
					row, err := strconv.Atoi(os.Args[3])
					if err == nil {
						p.LspTargetRow = row
					}
				}
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
