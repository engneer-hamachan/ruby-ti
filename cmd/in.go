package cmd

import (
	"os"
	"slices"
	"strconv"
	"strings"
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

	if hasFlag("--all-type") {
		flags.IsAllType = true
	}

	if hasFlag("--extends") {
		flags.IsExtends = true
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
	row := getTargetRow()
	if row > 0 {
		p.LspTargetRow = row
	}
}

func getTargetRow() int {
	for _, arg := range os.Args {
		if strings.HasPrefix(arg, "--row=") {
			if row, err := strconv.Atoi(arg[6:]); err == nil {
				return row
			}
		}
	}
	return 0
}

func ValidateArgs() {
	if hasFlag("--all-type") {
		return
	}

	if hasFlag("--extends") {
		return
	}

	if len(os.Args) == 1 {
		panic("want one argument!")
	}
}

func GetTargetFile() string {
	return os.Args[1]
}

func GetTargetClassName() string {
	for i, arg := range os.Args {
		if arg == "--extends" && i+1 < len(os.Args) {
			return os.Args[i+1]
		}
	}
	return ""
}
