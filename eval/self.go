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

	p.SetLastEvaluatedT(base.MakeSelf())

	return nil
}
