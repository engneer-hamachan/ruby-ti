package eval

import (
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

	if nextT == nil {
		return nil
	}

	err = e.Eval(p, ctx, nextT)
	if err != nil {
		return err
	}

	rightT := p.GetLastEvaluatedT()

	rangeT := base.MakeRange()

	switch leftT.GetType() {
	case rightT.GetType():
		rangeT.AppendArrayVariant(rightT)
	default:
		rangeT.AppendArrayVariant(*base.MakeUntyped())
	}

	p.SetLastEvaluatedT(rangeT)

	return nil
}
