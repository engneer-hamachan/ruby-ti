package eval

import (
	"ti/base"
	"ti/context"
	"ti/parser"
)

type OpenParentheses struct{}

func NewOpenParentheses() DynamicEvaluator {
	return &OpenParentheses{}
}

func init() {
	DynamicEvaluators["("] = NewOpenParentheses()
}

func (o *OpenParentheses) Evaluation(
	e *Evaluator,
	p *parser.Parser,
	ctx context.Context,
	t *base.T,
) (err error) {

	err = e.EvalToTargetToken(p, ctx, ")")
	if err != nil {
		return err
	}

	err = e.evalPriorityExp(p, ctx)
	if err != nil {
		return err
	}

	return nil
}
