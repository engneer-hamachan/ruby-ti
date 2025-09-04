package eval

import (
	"fmt"
	"ti/base"
	"ti/context"
	"ti/parser"
)

type DebugTypePrint struct{}

func NewDebugTypePrint() DynamicEvaluator {
	return &DebugTypePrint{}
}

func init() {
	DynamicEvaluators["dbtp"] = NewDebugTypePrint()
}

func (d *DebugTypePrint) Evaluation(
	e *Evaluator,
	p *parser.Parser,
	ctx context.Context,
	t *base.T,
) (err error) {

	row := p.ErrorRow

	nextT, err := p.Read()
	if err != nil {
		p.Fatal(ctx, err)
	}

	for {
		err = e.Eval(p, ctx, nextT)
		if err != nil {
			p.Fatal(ctx, err)
		}

		nextT, err = p.Read()
		if err != nil {
			p.Fatal(ctx, err)
		}

		if nextT.GetPower() == 0 {
			p.Unget()
			break
		}
	}

	if !ctx.IsCheckRound() {
		return nil
	}

	printT := p.GetLastEvaluatedT()

	p.ErrorRow = row

	p.Fatal(ctx, fmt.Errorf("%v", base.TypeToString(&printT)))

	return nil
}
