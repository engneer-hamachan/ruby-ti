package eval

import (
	"ti/base"
	"ti/context"
	"ti/parser"
)

type Return struct{}

func NewReturn() DynamicEvaluator {
	return &Return{}
}

func init() {
	DynamicEvaluators["return"] = NewReturn()
}

func (r *Return) Evaluation(
	e *Evaluator,
	p *parser.Parser,
	ctx context.Context,
	t *base.T,
) (err error) {

	nextT, err := p.Read()
	if err != nil {
		return err
	}

	p.StartParsingExpression()

	err = e.ContinuousEval(p, ctx, nextT, ",")
	if err != nil {
		return err
	}

	p.AppendLastReturnT()

	return nil
}
