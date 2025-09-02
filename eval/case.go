package eval

import (
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
			)
		}()
	}

	var caseTs []base.T
	var resultTs []base.T
	var isElse bool
	var isFirstWhen bool

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

		if nextT.IsTargetIdentifier("when") {
			if !isFirstWhen {
				isFirstWhen = true
			} else {
				resultTs = append(resultTs, p.GetLastEvaluatedT())
			}

			if objectT.IsIdentifierType() {
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
				)
			}

			continue
		}

		if nextT.IsTargetIdentifier("else") {
			isElse = true
			resultTs = append(resultTs, p.GetLastEvaluatedT())

			if evaluatedT.IsUnionType() && objectT.IsIdentifierType() {
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
				)
			}

			continue
		}

		err = e.Eval(p, ctx, nextT)
		if err != nil {
			return err
		}
	}

	if !isElse {
		resultTs = append(resultTs, *base.MakeNil())
	}

	resultT := base.MakeUnifiedT(resultTs)
	e.setLastEvaluatedT(p, ctx, resultT)

	return nil
}
