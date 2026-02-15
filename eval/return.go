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

	p.SetLastEvaluatedT(base.MakeUnknown())

	nextT, err := p.Read()
	if err != nil {
		return err
	}

	if nextT == nil {
		return nil
	}

	p.StartParsingExpression()

	err = e.ContinuousEval(p, ctx, nextT, ",")
	if err != nil {
		return err
	}

	nextT, err = p.Read()
	if err != nil {
		return err
	}

	if nextT.IsTargetIdentifier("[") {
		err = e.Eval(p, ctx, nextT)
		if err != nil {
			return err
		}

		p.AppendLastReturnT()

		return nil
	}

	p.Unget()
	p.AppendLastReturnT()

	return nil
}
