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
	var variable string

	if t.IsKeyIdentifier() {
		variable = t.ToString()[:len(t.ToString())-1]
	} else {
		variable = t.ToString()
	}

	base.SetValueT(
		ctx.GetFrame(),
		ctx.GetClass(),
		ctx.GetMethod(),
		variable,
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

	p.Unget()

	return nil
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

	// {name:, age:}
	case nextT.IsTargetIdentifier("{"):
		err := i.parseHash(p, ctx)
		if err != nil {
			return err
		}

	// name:
	case nextT.IsKeyIdentifier():
		variable := nextT.ToString()[:len(nextT.ToString())-1]
		nextT = base.MakeIdentifier(variable)

		i.parseVariable(ctx, nextT)

	// ^ ||  _
	case nextT.IsTargetPrefixIdentifier('^') || nextT.IsTargetIdentifier("_"):
		// nop

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

	// TODO:
	// 1. t.IsVariableIdentifier -> done
	// 2. implement bnf
	//   pattern :=
	//     literal
	//   | variable -> done
	//   | _ -> done
	//   | [pattern, ...] -> done
	//
	//   | { key: pattern, **pattern } -> done
	//   | Class -> done
	//   | Class[pattern, ...] -> done
	//   | Range
	//   | pattern | pattern -> done
	//   | pattern & pattern -> done
	//   | ^variable -> done
	//   | pattern if expr
	//   | pattern => variable -> done
	// 3. array inference
	// 4. hash inference
	nextT, err := p.Read()
	if err != nil {
		return err
	}

	i.parsePattern(p, ctx, nextT, true)

	p.SkipToTargetToken("\n")

	return nil
}
