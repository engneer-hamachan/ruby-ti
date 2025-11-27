package eval

import (
	"ti/base"
	"ti/context"
	"ti/parser"
)

type Self struct{}

func NewSelf() DynamicEvaluator {
	return &Self{}
}

func init() {
	DynamicEvaluators["self"] = NewSelf()
}

func (s *Self) Evaluation(
	e *Evaluator,
	p *parser.Parser,
	ctx context.Context,
	t *base.T,
) (err error) {

	var selfT *base.T
	switch ctx.IsDefineStatic {
	case true:
		selfT = base.MakeClass(ctx.GetClass())
	default:
		selfT = base.MakeSelf()
		selfT.DefinedFrame = ctx.GetFrame()
		selfT.DefinedClass = ctx.GetClass()
	}

	p.SetLastEvaluatedT(selfT)

	return nil
}
