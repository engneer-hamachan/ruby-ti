package eval

import (
	"fmt"
	"ti/base"
	"ti/context"
	"ti/parser"
)

type Ternary struct{}

func NewTernary() DynamicEvaluator {
	return &Ternary{}
}

func init() {
	DynamicEvaluators["?"] = NewTernary()
}

func (r *Ternary) Evaluation(
	e *Evaluator,
	p *parser.Parser,
	ctx context.Context,
	t *base.T,
) (err error) {

	var variants []base.T

	nextT, err := p.Read()
	if err != nil {
		return err
	}

	if nextT == nil {
		return nil
	}

	err = e.Eval(p, ctx, nextT)
	if err != nil {
		return err
	}

	variants = append(variants, p.GetLastEvaluatedT())

	nextT, err = p.Read()
	if err != nil {
		return err
	}

	if !nextT.IsTargetIdentifier(":") {
		return fmt.Errorf("syntax error")
	}

	nextT, err = p.Read()
	if err != nil {
		return err
	}

	err = e.Eval(p, ctx, nextT)
	if err != nil {
		return err
	}

	variants = append(variants, p.GetLastEvaluatedT())

	unionT := base.MakeUnion(variants)

	p.SetLastEvaluatedT(unionT.UnifyVariants())

	return nil
}
