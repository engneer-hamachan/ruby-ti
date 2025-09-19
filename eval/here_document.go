package eval

import (
	"ti/base"
	"ti/parser"
)

func evalHereDocument(
	p *parser.Parser,
	t *base.T,
) (err error) {

	target := t.ToString()[2:]

	isNewLine := false

	for {
		nextT, err := p.Read()
		if err != nil {
			return err
		}

		if nextT == nil {
			return nil
		}

		if nextT.IsNewLineIdentifier() {
			isNewLine = true
		}

		if nextT.ToString() == target {
			break
		}
	}

	if isNewLine {
		p.Row++
	}

	p.SetLastEvaluatedT(base.MakeAnyString())

	return nil
}
