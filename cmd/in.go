package cmd

import (
	"os"
	"slices"
	"strconv"
	"ti/parser"
)

func BuildFlags() *ExecuteFlags {
	flags := NewExecuteFlags()

	if hasFlag("-i") {
		flags.IsDefineInfo = true
	}

	if hasFlag("--define") {
		flags.IsDefineAllInfo = true
	}

	if hasFlag("--suggest") {
		flags.IsLsp = true
	}

	return flags
}

func ApplyParserFlags(p *parser.Parser) {
	if hasFlag("-d") {
		p.Debug = true
	}

	if hasFlag("--define") || hasFlag("--suggest") {
		applyTargetRow(p)
	}
}

func hasFlag(flag string) bool {
	return len(os.Args) > 0 && slices.Contains(os.Args, flag)
}

func applyTargetRow(p *parser.Parser) {
	if len(os.Args) > 3 {
		row, err := strconv.Atoi(os.Args[3])
		if err == nil {
			p.LspTargetRow = row
		}
	}
}

func ValidateArgs() {
	if len(os.Args) == 1 {
		panic("want one argument!")
	}
}

func GetTargetFile() string {
	return os.Args[1]
}
