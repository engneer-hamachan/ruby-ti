package eval

import (
	"ti/base"
	"ti/context"
	"ti/parser"
)

type Logical struct{}

func NewLogical() DynamicEvaluator {
	return &Logical{}
}

func init() {
	logical := NewLogical()
	DynamicEvaluators["&&"] = logical
	DynamicEvaluators["||"] = logical
	DynamicEvaluators["and"] = logical
	DynamicEvaluators["or"] = logical
}

func (r *Logical) Evaluation(
	e *Evaluator,
	p *parser.Parser,
	ctx context.Context,
	t *base.T,
) (err error) {

	var variants []base.T

	variants = append(variants, p.GetLastEvaluatedT())

	nextT, err := p.Read()
	if err != nil {
		return err
	}

	for {
		err = e.Eval(p, ctx, nextT)
		if err != nil {
			return err
		}

		nextT, err = p.Read()
		if err != nil {
			return err
		}

		if nextT == nil {
			break
		}

		if nextT.GetPower() == 0 && !nextT.IsTargetIdentifier("[") {
			p.Unget()
			break
		}
	}

	variants = append(variants, p.GetLastEvaluatedT())

	unionT := base.MakeUnifiedT(variants)

	p.SetLastEvaluatedT(unionT)

	return nil
}
