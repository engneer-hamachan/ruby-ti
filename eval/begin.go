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

	for {
		nextT, err := p.Read()
		if err != nil {
			return err
		}

		if nextT == nil {
			return nil
		}

		if nextT.IsEndIdentifier() {
			break
		}

		err = e.Eval(p, ctx, nextT)
		if err != nil {
			return err
		}
	}

	return nil
}
