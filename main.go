package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
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

	if len(base.TSignatures) > 0 && flags.IsSuggest {
		cmd.PrintSuggestionsForLsp(p)
	}

	if flags.IsExtends {
		cmd.PrintTargetClassExtends()
		os.Exit(0)
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

type TiIncludeConfig struct {
	Preload []string `json:"preload"`
}

func getPreloadFiles() []string {
	configPath := ".ti-include.json"

	data, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return []string{}
		}
		fmt.Fprintf(os.Stderr, "Error reading %s: %v\n", configPath, err)
		os.Exit(1)
	}

	var config TiIncludeConfig
	if err := json.Unmarshal(data, &config); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing %s: %v\n", configPath, err)
		os.Exit(1)
	}

	configDir, _ := os.Getwd()

	resolvedFiles := make([]string, len(config.Preload))
	for i, file := range config.Preload {
		if filepath.IsAbs(file) {
			resolvedFiles[i] = file
		} else {
			resolvedFiles[i] = filepath.Join(configDir, file)
		}
	}

	return resolvedFiles
}

func preload(round string, flags *cmd.ExecuteFlags) {
	for _, preloadFile := range getPreloadFiles() {
		if fp, err := os.Open(preloadFile); err == nil {
			br := bufio.NewReader(fp)
			p := getParser(br, preloadFile)
			cmd.ApplyParserFlags(&p)
			loop(p, flags, round)
			fp.Close()
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
			preload(round, flags)

			file = cmd.GetTargetFile()
			fp, _ := os.Open(file)
			br = bufio.NewReader(fp)

			p := getParser(br, file)
			cmd.ApplyParserFlags(&p)

			cleanSimpleIdentifires()

			loop(p, flags, round)
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
