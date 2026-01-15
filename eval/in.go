package eval

import (
	"ti/base"
	"ti/builtin"
	"ti/context"
	"ti/parser"
)

type In struct {
	caseTargetT base.T
}

func NewIn() DynamicEvaluator {
	return &In{}
}

func init() {
	DynamicEvaluators["in"] = NewIn()
}

func (i *In) parseVariable(ctx context.Context, t *base.T) {
	base.SetValueT(
		ctx.GetFrame(),
		ctx.GetClass(),
		ctx.GetMethod(),
		t.ToString(),
		base.MakeUntyped(),
		ctx.IsDefineStatic,
	)
}

func (i *In) parseHash(
	p *parser.Parser,
	ctx context.Context,
) error {

	for {
		nextT, err := p.Read()
		if err != nil {
			return err
		}

		if nextT.IsTargetIdentifier("}") {
			break
		}

		err = i.parsePattern(p, ctx, nextT, false)
		if err != nil {
			return err
		}
	}

	p.SetLastEvaluatedT(base.MakeUntyped())

	return nil
}

func (i *In) parseArray(
	p *parser.Parser,
	ctx context.Context,
) error {

	ct := 1
	arrayVariants := i.caseTargetT.GetVariants()
	variantsLength := len(arrayVariants)

	for {
		nextT, err := p.Read()
		if err != nil {
			return err
		}

		if nextT.IsVariableIdentifier() {
			var objectT *base.T

			if ct <= variantsLength {
				objectT = arrayVariants[ct-1].DeepCopy()
			} else {
				objectT = i.caseTargetT.UnifyVariants()
			}

			base.SetValueT(
				ctx.GetFrame(),
				ctx.GetClass(),
				ctx.GetMethod(),
				nextT.ToString(),
				objectT,
				ctx.IsDefineStatic,
			)
		}

		if nextT.IsAsteriskPrefix() {
			base.SetValueT(
				ctx.GetFrame(),
				ctx.GetClass(),
				ctx.GetMethod(),
				nextT.ToString()[1:],
				base.MakeUntyped(),
				ctx.IsDefineStatic,
			)
		}

		if nextT.IsClassType() {
			err = i.parsePattern(p, ctx, nextT, false)
			if err != nil {
				return err
			}
		}

		if nextT.IsCommaIdentifier() {
			continue
		}

		if nextT.IsTargetIdentifier("]") {
			break
		}

		ct++
	}

	p.SetLastEvaluatedT(base.MakeUntyped())

	return nil
}

func (i *In) parseParentheses(p *parser.Parser, ctx context.Context) error {
	p.SetLastEvaluatedT(base.MakeUntyped())

	for {
		nextT, err := p.Read()
		if err != nil {
			return err
		}

		if nextT.IsVariableIdentifier() {
			i.parseVariable(ctx, nextT)
		}

		if nextT.IsTargetIdentifier(")") {
			break
		}
	}

	return nil
}

func (i *In) parseClass(
	p *parser.Parser,
	ctx context.Context,
	t *base.T,
) error {

	objectT := builtin.ConvertToBuiltinT(t.ToString())
	p.SetLastEvaluatedT(objectT.DeepCopy())

	// [
	nextT, err := p.Read()
	if err != nil {
		return err
	}

	if nextT.IsTargetIdentifier("[") {
		err := i.parseArray(p, ctx)
		if err != nil {
			return err
		}

		return nil
	}

	if nextT.IsTargetIdentifier("(") {
		err := i.parseParentheses(p, ctx)
		if err != nil {
			return err
		}

		return nil
	}

	p.Unget()

	return nil
}

func containKey(ctx context.Context, variable string, variants []base.T) bool {
	for _, variant := range variants {
		if variant.GetRemovePrefixKey() == variable {
			base.SetValueT(
				ctx.GetFrame(),
				ctx.GetClass(),
				ctx.GetMethod(),
				variable,
				variant.GetKeyValue(),
				ctx.IsDefineStatic,
			)

			return true
		}

		if variant.GetKeyValue().IsHashType() {
			return containKey(ctx, variable, variant.GetKeyValue().GetVariants())
		}
	}

	return false
}

func (i *In) parsePattern(
	p *parser.Parser,
	ctx context.Context,
	nextT *base.T,
	isFirstToken bool,
) error {

	switch {
	// String
	case nextT.IsClassType():
		err := i.parseClass(p, ctx, nextT)
		if err != nil {
			return err
		}

	// [x, y]
	case nextT.IsTargetIdentifier("["):
		err := i.parseArray(p, ctx)
		if err != nil {
			return err
		}

	// ()
	case nextT.IsOpenParentheses():
		err := i.parseParentheses(p, ctx)
		if err != nil {
			return err
		}

	// {name:, age:}
	case nextT.IsTargetIdentifier("{"):
		err := i.parseHash(p, ctx)
		if err != nil {
			return err
		}

	// **y
	case nextT.IsDoubleAsteriskPrefix():
		variable := nextT.ToString()[2:]

		base.SetValueT(
			ctx.GetFrame(),
			ctx.GetClass(),
			ctx.GetMethod(),
			variable,
			base.MakeAnyHash(),
			ctx.IsDefineStatic,
		)

	case nextT.IsTargetIdentifier("^"):
		p.SetLastEvaluatedT(base.MakeUntyped())

		nextT, err := p.Read()
		if err != nil {
			return err
		}

		err = i.parsePattern(p, ctx, nextT, false)
		if err != nil {
			return err
		}

	// name:
	case nextT.IsKeyIdentifier():
		variable := nextT.ToString()[:len(nextT.ToString())-1]

		isContain := containKey(ctx, variable, i.caseTargetT.GetVariants())
		if isContain {
			break
		}

		nextT = base.MakeIdentifier(variable)
		i.parseVariable(ctx, nextT)

	// ^ ||  _
	case nextT.IsTargetPrefixIdentifier('^') || nextT.IsTargetIdentifier("_"):
	//// nop

	// x
	case nextT.IsVariableIdentifier():
		i.parseVariable(ctx, nextT)
	}

	// pattern => x
	nextT, err := p.Read()
	if err != nil {
		return err
	}

	switch nextT.ToString() {
	case "=>":
		variableT, err := p.Read()
		if err != nil {
			return err
		}

		evaluatedT := p.GetLastEvaluatedT()

		base.SetValueT(
			ctx.GetFrame(),
			ctx.GetClass(),
			ctx.GetMethod(),
			variableT.ToString(),
			&evaluatedT,
			ctx.IsDefineStatic,
		)

	case ",", "&", "|":
		nextT, err := p.Read()
		if err != nil {
			return err
		}

		err = i.parsePattern(p, ctx, nextT, false)
		if err != nil {
			return err
		}

	default:
		if isFirstToken {
			evaluatedT := p.GetLastEvaluatedT()

			base.SetValueT(
				ctx.GetFrame(),
				ctx.GetClass(),
				ctx.GetMethod(),
				i.caseTargetT.GetBeforeEvaluateCode(),
				&evaluatedT,
				ctx.IsDefineStatic,
			)
		}
		p.Unget()
	}

	return nil
}

func (i *In) Evaluation(
	e *Evaluator,
	p *parser.Parser,
	ctx context.Context,
	t *base.T,
) (err error) {

	i.caseTargetT = p.GetLastEvaluatedT()

	for {
		nextT, err := p.Read()
		if err != nil {
			return err
		}

		if nextT.IsNewLineIdentifier() {
			break
		}

		i.parsePattern(p, ctx, nextT, true)
	}

	return nil
}
