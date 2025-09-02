package eval

import (
	"fmt"
	"ti/base"
	"ti/context"
	"ti/parser"
)

type Range struct{}

func NewRange() DynamicEvaluator {
	return &Range{}
}

func init() {
	r := NewRange()
	DynamicEvaluators[".."] = r
	DynamicEvaluators["..."] = r
}

func (r *Range) Evaluation(
	e *Evaluator,
	p *parser.Parser,
	ctx context.Context,
	t *base.T,
) (err error) {

	leftT := p.GetLastEvaluatedT()

	nextT, err := p.Read()
	if err != nil {
		return err
	}

	err = e.Eval(p, ctx, nextT)
	if err != nil {
		return err
	}

	rightT := p.GetLastEvaluatedT()

	if leftT.GetType() != rightT.GetType() {
		return fmt.Errorf("not equal left right term")
	}

	rangeT := base.MakeRange()
	rangeT.AppendArrayVariant(rightT)

	p.SetLastEvaluatedT(rangeT)

	return nil
}
