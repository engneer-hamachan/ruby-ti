package eval

import (
	"strings"
	"ti/base"
	"ti/parser"
)

func evalHereDocument(
	p *parser.Parser,
	t *base.T,
) (err error) {

	target := t.ToString()[2:]

	for {
		nextT, err := p.Read()
		if err != nil {
			return err
		}

		if nextT == nil {
			return nil
		}

		if strings.Contains(target, nextT.ToString()) {
			break
		}
	}

	p.SetLastEvaluatedT(base.MakeAnyString())

	return nil
}
