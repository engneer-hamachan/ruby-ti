package eval

import (
	"ti/parser"
)

func skipMultilineComment(p *parser.Parser) (err error) {
	for {
		t, err := p.Read()
		if err != nil {
			return err
		}

		if t == nil {
			return nil
		}

		nextT, err := p.Read()
		if err != nil {
			return err
		}

		if t.IsEqualIdentifier() && nextT.IsTargetIdentifier("end") {
			return nil
		}
	}
}
