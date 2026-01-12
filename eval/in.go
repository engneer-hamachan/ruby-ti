package eval

import (
	"slices"
	"ti/base"
	"ti/context"
	"ti/parser"
)

type In struct{}

func NewIn() DynamicEvaluator {
	return &In{}
}

func init() {
	DynamicEvaluators["in"] = NewIn()
}

func (i *In) Evaluation(
	e *Evaluator,
	p *parser.Parser,
	ctx context.Context,
	t *base.T,
) (err error) {

	skipTarget := []string{"[", "]", "{", "}", ",", "\n"}

	// TODO:
	// 1. t.IsVariableIdentifier
	// 2. implement bnf
	//   pattern :=
	//     literal
	//   | variable
	//   | _
	//   | [pattern, ...]
	//
	//   | { key: pattern, **pattern }
	//   | Class
	//   | Class[pattern, ...]
	//   | Range
	//   | pattern | pattern
	//   | pattern & pattern
	//   | ^variable
	//   | pattern if expr
	//   | pattern => variable
	// 3. array inference
	// 4. hash inference

	// parsePattern
	nextT, err := p.Read()
	if err != nil {
		return err
	}

	// in String => x
	if nextT.IsClassType() {
		objectT := base.MakeObject(nextT.ToString())

		// =>
		nextT, err := p.Read()
		if err != nil {
			return err
		}

		if nextT.IsTargetIdentifier("=>") {
			// x
			nextT, err := p.Read()
			if err != nil {
				return err
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
	}

	// [x, y]
	if nextT.IsTargetIdentifier("[") {
		for {
			nextT, err := p.Read()
			if err != nil {
				return err
			}

			if nextT.GetPower() == 0 && nextT.IsUnknownType() && !slices.Contains(skipTarget, nextT.ToString()) {
				base.SetValueT(
					ctx.GetFrame(),
					ctx.GetClass(),
					ctx.GetMethod(),
					nextT.ToString(),
					base.MakeUntyped(),
					ctx.IsDefineStatic,
				)
			}

			if nextT.IsTargetIdentifier("]") {
				break
			}
		}
	}

	// {name:, age:}
	if nextT.IsTargetIdentifier("{") {
		for {
			nextT, err := p.Read()
			if err != nil {
				return err
			}

			if nextT.GetPower() == 0 && nextT.IsUnknownType() && !slices.Contains(skipTarget, nextT.ToString()) {
				var variable string

				if nextT.IsKeyIdentifier() {
					variable = nextT.ToString()[:len(nextT.ToString())-1]
				} else {
					variable = nextT.ToString()
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

			if nextT.IsTargetIdentifier("}") {
				break
			}
		}
	}

	// x
	if nextT.IsIdentifierType() && !slices.Contains(skipTarget, nextT.ToString()) && nextT.GetPower() == 0 {
		var variable string

		if nextT.IsKeyIdentifier() {
			variable = nextT.ToString()[:len(nextT.ToString())-1]
		} else {
			variable = nextT.ToString()
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

	p.SkipToTargetToken("\n")

	return nil
}
