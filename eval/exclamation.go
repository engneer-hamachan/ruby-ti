package eval

import (
	"ti/base"
	"ti/context"
	"ti/parser"
)

type Exclamation struct{}

func NewExclamation() DynamicEvaluator {
	return &Exclamation{}
}

func init() {
	DynamicEvaluators["!"] = NewExclamation()
}

func (e *Exclamation) Evaluation(
	ev *Evaluator,
	p *parser.Parser,
	ctx context.Context,
	t *base.T,
) (err error) {

	t, err = p.Read()
	if err != nil {
		return err
	}

	err = ev.Eval(p, ctx, t)
	if err != nil {
		return err
	}

	p.SetLastEvaluatedT(base.MakeBool())

	return nil
}
