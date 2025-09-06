package eval

import (
	"fmt"
	"ti/base"
	"ti/context"
	"ti/parser"
)

type Bind struct{}

func NewBind() DynamicEvaluator {
	return &Bind{}
}

func init() {
	bind := NewBind()
	DynamicEvaluators["="] = bind
	DynamicEvaluators["||="] = bind
}

func (b *Bind) handleScalarAsigntment(
	e *Evaluator,
	p *parser.Parser,
	ctx context.Context,
) (err error) {

	leftT := p.GetLastEvaluatedTPointer().(*base.T)

	ctx.IsBind = true

	nextT, err := p.Read()
	if err != nil {
		return err
	}

	for {
		err = e.ContinuousEval(p, ctx, nextT, ",")
		if err != nil {
			return err
		}

		nextT, err = p.Read()
		if err != nil {
			return err
		}

		if nextT.GetPower() == 0 {
			p.Unget()
			break
		}
	}

	rightT := p.GetLastEvaluatedT()

	if leftT.HasDefault() {
		return nil
	}

	if leftT.IsReadOnly() && leftT.GetBeforeEvaluateCode()[0] != '@' {
		return fmt.Errorf("%s is read only", leftT.GetBeforeEvaluateCode())
	}

	if leftT.IsReadOnly() {
		rightT.EnableReadOnly()
	}

	rightT.SetHasDefault(ctx.IsDefineArg)

	*leftT = rightT

	p.SetLastEvaluatedT(leftT)

	return nil
}

func (b *Bind) handleMultipleToScalarAsigntment(
	ctx context.Context,
	leftTs []*base.T,
	rightT *base.T,
) error {

	var idx int

	switch rightT.IsArrayType() {
	case true:
		for {
			if (idx + 1) > len(leftTs) {
				break
			}

			if (idx + 1) > len(rightT.GetVariants()) {
				*leftTs[idx] = *base.MakeNil()
				idx++
				continue
			}

			if leftTs[idx].HasDefault() {
				return nil
			}

			if leftTs[idx].IsReadOnly() && leftTs[idx].GetBeforeEvaluateCode()[0] != '@' {
				return fmt.Errorf("%s is read only", leftTs[idx].GetBeforeEvaluateCode())
			}

			unionT := rightT.UnifyVariants()
			unionT.SetHasDefault(ctx.IsDefineArg)

			if leftTs[idx].IsReadOnly() {
				unionT.EnableReadOnly()
			}

			*leftTs[idx] = *unionT

			idx++
		}

	default:
		if leftTs[0].HasDefault() {
			return nil
		}

		if leftTs[0].IsReadOnly() && leftTs[0].GetBeforeEvaluateCode()[0] != '@' {
			return fmt.Errorf("%s is read only", leftTs[0].GetBeforeEvaluateCode())
		}

		rightT.SetHasDefault(ctx.IsDefineArg)

		if leftTs[idx].IsReadOnly() {
			rightT.EnableReadOnly()
		}

		rightT.SetHasDefault(ctx.IsDefineArg)
		*leftTs[0] = *rightT
	}

	return nil
}

func (b *Bind) handleMultipleToMultipleAsigntment(
	leftTs []*base.T,
	rightTs *base.T,
) error {

	var idx int
	for {
		if (idx + 1) > len(leftTs) {
			break
		}

		if (idx + 1) > len(rightTs.GetVariants()) {
			*leftTs[idx] = *base.MakeNil()
			idx++
			continue
		}

		if leftTs[idx].HasDefault() {
			return nil
		}

		if leftTs[idx].IsReadOnly() && leftTs[idx].GetBeforeEvaluateCode()[0] != '@' {
			return fmt.Errorf("%s is read only", leftTs[idx].GetBeforeEvaluateCode())
		}

		if leftTs[idx].IsReadOnly() {
			rightTs.GetVariants()[idx].EnableReadOnly()
		}

		*leftTs[idx] = rightTs.GetVariants()[idx]

		idx++
	}

	return nil
}

func (b *Bind) handleMultipleAsigntment(
	e *Evaluator,
	p *parser.Parser,
	ctx context.Context,
) (err error) {

	leftTs := p.GetLastEvaluatedTPointer().([]*base.T)

	nextT, err := p.Read()
	if err != nil {
		return err
	}

	err = e.ContinuousEval(p, ctx, nextT, ",")
	if err != nil {
		return err
	}

	rightT := p.GetLastEvaluatedT()

	switch rightT.IsArrayType() {
	case true:
		return b.handleMultipleToMultipleAsigntment(leftTs, &rightT)

	default:
		return b.handleMultipleToScalarAsigntment(ctx, leftTs, &rightT)
	}
}

func (b *Bind) Evaluation(
	e *Evaluator,
	p *parser.Parser,
	ctx context.Context,
	t *base.T,
) (err error) {

	p.EndParsingExpression()

	some := p.GetLastEvaluatedTPointer()

	if err := p.SkipNewline(); err != nil {
		return err
	}

	switch some.(type) {
	case *base.T:
		return b.handleScalarAsigntment(e, p, ctx)

	case []*base.T:
		return b.handleMultipleAsigntment(e, p, ctx)

	default:
		return fmt.Errorf("syntax error")
	}
}
