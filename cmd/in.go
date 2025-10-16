package cmd

import (
	"os"
	"slices"
	"strconv"
	"ti/parser"
)

func ParseArgs(p *parser.Parser) {
	if len(os.Args) > 0 && slices.Contains(os.Args, "-d") {
		p.Debug = true
	}

	if len(os.Args) > 0 && slices.Contains(os.Args, "-i") {
		p.IsDefineInfo = true
	}

	if len(os.Args) > 0 && slices.Contains(os.Args, "--define") {
		p.IsDefineAllInfo = true
		if len(os.Args) > 3 {
			row, err := strconv.Atoi(os.Args[3])
			if err == nil {
				p.LspTargetRow = row
			}
		}
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
}

func ValidateArgs() {
	if len(os.Args) == 1 {
		panic("want one argument!")
	}
}

func GetTargetFile() string {
	return os.Args[1]
}
