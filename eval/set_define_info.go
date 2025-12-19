package eval

import (
	"fmt"
	"ti/base"
	"ti/parser"
)

func setDefineInfos(p *parser.Parser, defineRow int) {
	var hint string

	hint += "@"
	hint += p.FileName + ":::"
	hint += fmt.Sprintf("%d", defineRow)
	hint += ":::"

	hint += "bind: "
	t := p.GetLastEvaluatedT()
	hint += base.TypeToString(&t)

	p.DefineInfos = append(p.DefineInfos, hint)
}
