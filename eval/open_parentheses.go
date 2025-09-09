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

	for {
		nextT, err := p.Read()
		if err != nil {
			return err
		}

		if nextT == nil {
			break
		}

		if nextT.IsCloseParentheses() {
			break
		}

		err = e.Eval(p, ctx, nextT)
		if err != nil {
			return err
		}
	}

	err = e.evalPriorityExp(p, ctx)
	if err != nil {
		return err
	}

	return nil
}
