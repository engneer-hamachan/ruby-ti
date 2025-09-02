package eval

import (
	"fmt"
	"ti/base"
	"ti/context"
	"ti/parser"
)

type DebugPrint struct{}

func NewDebugPrint() DynamicEvaluator {
	return &DebugPrint{}
}

func init() {
	DynamicEvaluators["dbp"] = NewDebugPrint()
}

func (d *DebugPrint) Evaluation(
	e *Evaluator,
	p *parser.Parser,
	ctx context.Context,
	t *base.T,
) (err error) {

	nextT, err := p.Read()
	if err != nil {
		return err
	}

	err = e.Eval(p, ctx, nextT)
	if err != nil {
		return err
	}

	return fmt.Errorf("%v", p.GetLastEvaluatedT())
}
