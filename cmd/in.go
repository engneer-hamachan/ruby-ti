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
		flags.IsSuggest = true
	}

	if hasFlag("--all-type") {
		flags.IsAllType = true
	}

	if hasFlag("--extends") {
		flags.IsExtends = true
	}

	if hasFlag("--hover") {
		flags.IsHover = true
	}

	return flags
}

func ApplyParserFlags(p *parser.Parser) {
	if hasFlag("-d") {
		p.Debug = true
	}

	if hasFlag("--define") || hasFlag("--suggest") || hasFlag("--hover") {
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

func getTargetClass() string {
	for _, arg := range os.Args {
		if strings.HasPrefix(arg, "--class=") {
			return arg[8:]
		}
	}
	return ""
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
