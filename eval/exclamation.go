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

	p.LastCallT = t.DeepCopy()

	t, err = p.Read()
	if err != nil {
		return err
	}

	err = ev.Eval(p, ctx, t)
	if err != nil {
		return err
	}

	for {
		t, err := p.Read()
		if err != nil {
			return err
		}

		if t == nil {
			return nil
		}

		if t.IsCommaIdentifier() {
			continue
		}

		if t.GetPower() < 95 {
			p.Unget()
			break
		}

		err = ev.Eval(p, ctx, t)
		if err != nil {
			return err
		}
	}

	p.SetLastEvaluatedT(base.MakeBool())

	return nil
}
