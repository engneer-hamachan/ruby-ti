package eval

import (
	"ti/base"
	"ti/context"
	"ti/parser"
)

type While struct{}

func NewWhile() DynamicEvaluator {
	return &While{}
}

func init() {
	while := NewWhile()
	DynamicEvaluators["while"] = while
	DynamicEvaluators["until"] = while
	DynamicEvaluators["for"] = while
}

func (d *While) Evaluation(
	e *Evaluator,
	p *parser.Parser,
	ctx context.Context,
	t *base.T,
) (err error) {

	forTargetT := base.MakeUntyped()

	if t.IsTargetIdentifier("for") {
		for {
			nextT, err := p.Read()
			if err != nil {
				return err
			}

			if nextT.IsTargetIdentifier("in") {
				nextT, err := p.Read()
				if err != nil {
					return err
				}

				err = e.Eval(p, ctx, nextT)
				if err != nil {
					return err
				}

				lastEvaluatedT := p.GetLastEvaluatedT()
				if lastEvaluatedT.IsArrayType() || lastEvaluatedT.IsHashType() || lastEvaluatedT.IsRangeType() {
					forTargetT = &lastEvaluatedT
				}

				break
			}

			if nextT.IsCommaIdentifier() {
				continue
			}

			if !nextT.IsIdentifierType() {
				continue
			}

			base.SetValueT(
				ctx.GetFrame(),
				ctx.GetClass(),
				ctx.GetMethod(),
				nextT.ToString(),
				base.MakeUntyped(),
				ctx.IsDefineStatic,
			)
		}
	}

	isEatedNewlineToken := false
	isBreak := false

	for {
		nextT, err := p.Read()
		if err != nil {
			return err
		}

		if nextT == nil {
			break
		}

		if nextT.IsNewLineIdentifier() {
			isEatedNewlineToken = true
		}

		if nextT.IsTargetIdentifier("do") && !isEatedNewlineToken {
			continue
		}

		if nextT.IsTargetIdentifier("{") && !isEatedNewlineToken {
			continue
		}

		if nextT.IsTargetIdentifier("end") || nextT.IsTargetIdentifier("}") {
			break
		}

		if nextT.IsTargetIdentifier("break") {
			isBreak = true
			continue
		}

		err = e.Eval(p, ctx, nextT)
		if err != nil && !ctx.IsCollectRound() {
			return err
		}
	}

	if t.IsTargetIdentifier("for") && !isBreak {
		p.SetLastEvaluatedT(forTargetT)
		return nil
	}

	if t.IsTargetIdentifier("for") && isBreak {
		p.SetLastEvaluatedT(base.MakeUntyped())
		return nil
	}

	p.SetLastEvaluatedT(base.MakeNil())

	return nil
}
