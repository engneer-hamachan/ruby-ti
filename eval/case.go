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
	var isFirstWhen bool

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
		case "when":
			switch isFirstWhen {
			case true:
				resultTs = append(resultTs, p.GetLastEvaluatedT())
			case false:
				isFirstWhen = true
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
