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
		err = e.Eval(p, ctx, nextT)
		if err != nil {
			return err
		}

		nextT, err = p.Read()
		if err != nil {
			return err
		}

		if nextT == nil {
			return nil
		}

		if nextT.IsCommaIdentifier() {
			continue
		}

		if nextT.GetPower() == 0 {
			p.Unget()
			break
		}
	}

	rightT := p.GetLastEvaluatedT()

	// TODO: a[0]
	nextT, err = p.Read()
	if err != nil {
		return err
	}
	switch nextT.IsTargetIdentifier("[") {
	case true:
		e.Eval(p, ctx, nextT)
		rightT = p.GetLastEvaluatedT()
	default:
		p.Unget()
	}
	// TODO: end

	if leftT.HasDefault() {
		return nil
	}

	if leftT.IsReadOnly() && !leftT.IsBeforeEvaluateAtmarkPrefix() {
		return fmt.Errorf("%s is read only", leftT.GetBeforeEvaluateCode())
	}

	rightT.DisableReadOnly()

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

	var leftIdx int
	var rightIdx int

	rightVariants := rightTs.GetVariants()
	rightLen := len(rightTs.GetVariants())

	for {
		if (leftIdx + 1) > len(leftTs) {
			break
		}

		if rightIdx >= rightLen {
			*leftTs[leftIdx] = *base.MakeNil()

			leftIdx++
			rightIdx++

			continue
		}

		if leftTs[leftIdx].HasDefault() {
			return nil
		}

		if leftTs[leftIdx].IsReadOnly() &&
			!leftTs[leftIdx].IsBeforeEvaluateAtmarkPrefix() {

			return fmt.Errorf(
				"%s is read only",
				leftTs[leftIdx].GetBeforeEvaluateCode(),
			)
		}

		if leftTs[leftIdx].IsReadOnly() {
			rightVariants[rightIdx].EnableReadOnly()
		}

		// *x, y = 1, 2, 3
		if leftTs[leftIdx].IsBeforeEvaluateAsteriskPrefix() {
			arrayT := base.MakeAnyArray()

			for {
				if rightIdx > (rightLen - len(leftTs[leftIdx:])) {
					break
				}

				arrayT.AppendArrayVariant(rightVariants[rightIdx])

				rightIdx++
			}

			*leftTs[leftIdx] = *arrayT

			leftIdx++

			continue
		}

		*leftTs[leftIdx] = rightVariants[rightIdx]

		leftIdx++
		rightIdx++
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

func (b *Bind) setDefineInfos(p *parser.Parser, defineRow int) {
	var hint string

	hint += "@"
	hint += p.FileName + ":::"
	hint += fmt.Sprintf("%d", defineRow)
	hint += ":::"

	hint += "bind: "
	t := p.GetLastEvaluatedT()
	hint += base.TypeToString(&t)

	p.DefineInfos = append(p.DefineInfos, hint)
}

func (b *Bind) Evaluation(
	e *Evaluator,
	p *parser.Parser,
	ctx context.Context,
	t *base.T,
) (err error) {

	defer b.setDefineInfos(p, p.ErrorRow)

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
