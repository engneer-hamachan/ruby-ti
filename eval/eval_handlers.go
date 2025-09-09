package eval

import (
	"ti/base"
	"ti/context"
	"ti/eval/method_evaluator"
	"ti/parser"
)

func (e *Evaluator) handleEvaluateMethod(
	p *parser.Parser,
	ctx context.Context,
	objectT *base.T,
	methodT *base.T,
	isAmpersand bool,
) error {

	methodEvaluator :=
		method_evaluator.NewMethodEvaluator(
			e,
			p,
			ctx,
			objectT,
			methodT,
			isAmpersand,
		)

	return methodEvaluator.Evaluation()
}

func (e *Evaluator) handleDynamicEvaluator(
	p *parser.Parser,
	ctx context.Context,
	t *base.T,
) (bool, error) {

	id := t.ToString()

	dynamicEvalutor, ok := DynamicEvaluators[id]
	if ok {
		return true, dynamicEvalutor.Evaluation(e, p, ctx, t)
	}

	return false, nil
}

func (e *Evaluator) handleIdentifier(
	p *parser.Parser,
	ctx context.Context,
	t *base.T,
) {

	id := t.ToString()

	var valueT *base.T

	// <<EOF
	if len(id) > 2 && id[:2] == "<<" {
		evalHereDocument(p, t)
		return
	}

	switch rune(id[0]) {
	case '@':
		valueT =
			base.GetInstanceValueT(
				ctx.GetFrame(),
				ctx.GetClass(),
				id[1:],
			)

		if valueT == nil {
			identifierT := base.MakeIdentifier(id)

			base.SetInstanceValueT(
				ctx.GetFrame(),
				ctx.GetClass(),
				id[1:],
				identifierT,
			)

			e.setLastEvaluatedT(p, ctx, identifierT)

			return
		}

	default:
		valueT = base.GetValueT(ctx.GetFrame(), ctx.GetClass(), ctx.GetMethod(), id)
	}

	if valueT != nil {
		valueT.SetBeforeEvaluateCode(id)

		e.setLastEvaluatedT(p, ctx, valueT)

		return
	}

	identifierT := base.MakeIdentifier(id)

	base.SetValueT(
		ctx.GetFrame(),
		ctx.GetClass(),
		ctx.GetMethod(),
		id,
		identifierT,
	)

	e.setLastEvaluatedT(p, ctx, identifierT)
}

func (e *Evaluator) handleConstEvaluation(
	p *parser.Parser,
	ctx context.Context,
	t *base.T,
) error {

	p.Unget()

	valueT := base.GetConstValueT(ctx.GetFrame(), ctx.GetClass(), t.ToString())
	if valueT != nil {
		p.SetLastEvaluatedT(valueT)
		return nil
	}

	identifierT := base.MakeIdentifier(t.ToString())

	base.SetConstValueT(ctx.GetFrame(), ctx.GetClass(), t.ToString(), identifierT)

	e.setLastEvaluatedT(p, ctx, identifierT)
	return nil
}
