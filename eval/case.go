package eval

import (
	"fmt"
	"slices"
	"ti/base"
	"ti/context"
	"ti/parser"
)

type Case struct{}

func NewCase() DynamicEvaluator {
	return &Case{}
}

func init() {
	DynamicEvaluators["case"] = NewCase()
}

func (c *Case) Evaluation(
	e *Evaluator,
	p *parser.Parser,
	ctx context.Context,
	t *base.T,
) (err error) {

	objectT, err := p.Read()
	if err != nil {
		return err
	}

	err = e.Eval(p, ctx, objectT)
	if err != nil {
		return err
	}

	evaluatedT := p.GetLastEvaluatedT()

	if objectT.IsIdentifierType() {
		defer func() {
			base.SetValueT(
				ctx.GetFrame(),
				ctx.GetClass(),
				ctx.GetMethod(),
				objectT.ToString(),
				&evaluatedT,
				ctx.IsDefineStatic,
			)
		}()
	}

	var caseTs []base.T
	var isFirstBranch bool

	resultTs := []base.T{*base.MakeNil()}

	for {
		nextT, err := p.Read()
		if err != nil {
			return err
		}

		if nextT == nil {
			break
		}

		if nextT.IsEndIdentifier() {
			if !p.IsParsingExpression() {
				resultTs = append(resultTs, p.GetLastEvaluatedT())
			}

			break
		}

		switch nextT.ToString() {
		case "in":
			switch isFirstBranch {
			case true:
				resultTs = append(resultTs, p.GetLastEvaluatedT())
			case false:
				isFirstBranch = true
			}

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

			if nextT.IsTargetIdentifier("[") {
				for {
					nextT, err := p.Read()
					if err != nil {
						return err
					}

					if nextT.GetPower() == 0 && nextT.IsUnknownType() {
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

			if nextT.IsTargetIdentifier("{") {
				skipTarget := []string{"}", ","}

				for {
					nextT, err := p.Read()
					if err != nil {
						return err
					}

					if nextT.GetPower() == 0 && nextT.IsUnknownType() && !slices.Contains(skipTarget, nextT.ToString()) {
						var variable string

						if nextT.IsKeyIdentifier() {
							fmt.Println(nextT.ToString())
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

			p.SkipToTargetToken("\n")

		case "when":
			switch isFirstBranch {
			case true:
				resultTs = append(resultTs, p.GetLastEvaluatedT())
			case false:
				isFirstBranch = true
			}

			if !objectT.IsIdentifierType() {
				continue
			}

			caseT, err := p.Read()
			if err != nil {
				return err
			}

			err = e.Eval(p, ctx, caseT)
			if err != nil {
				return err
			}

			evaluatedT := p.GetLastEvaluatedT()
			caseTs = append(caseTs, evaluatedT)

			base.SetValueT(
				ctx.GetFrame(),
				ctx.GetClass(),
				ctx.GetMethod(),
				objectT.ToString(),
				&evaluatedT,
				ctx.IsDefineStatic,
			)

		case "else":
			resultTs = resultTs[1:]
			resultTs = append(resultTs, p.GetLastEvaluatedT())

			if !evaluatedT.IsUnionType() || !objectT.IsIdentifierType() {
				continue
			}

			unionVariants := evaluatedT.GetVariants()

			var newUnionVariants []base.T

			for _, variant := range unionVariants {
				isContain := false
				for _, caseT := range caseTs {
					if caseT.GetType() == variant.GetType() {
						isContain = true
						break
					}
				}

				if isContain {
					continue
				}

				newUnionVariants = append(newUnionVariants, variant)
			}

			unionT := base.MakeUnion(newUnionVariants)
			base.SetValueT(
				ctx.GetFrame(),
				ctx.GetClass(),
				ctx.GetMethod(),
				objectT.ToString(),
				unionT.UnifyVariants(),
				ctx.IsDefineStatic,
			)
		}

		err = e.Eval(p, ctx, nextT)
		if err != nil {
			return err
		}
	}

	resultT := base.MakeUnifiedT(resultTs)
	e.setLastEvaluatedT(p, ctx, resultT)

	return nil
}
