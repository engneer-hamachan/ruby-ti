package eval

import (
	"ti/base"
	"ti/context"
	"ti/parser"
)

type Begin struct{}

func NewBegin() DynamicEvaluator {
	return &Begin{}
}

func init() {
	DynamicEvaluators["begin"] = NewBegin()
}

func (r *Begin) Evaluation(
	e *Evaluator,
	p *parser.Parser,
	ctx context.Context,
	t *base.T,
) (err error) {

	err = e.EvalToTargetToken(p, ctx, "end")
	if err != nil {
		return err
	}

	return nil
}
