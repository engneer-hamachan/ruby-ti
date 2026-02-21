package method_evaluator

import (
	"strconv"
	"ti/base"
	"ti/context"
	"ti/parser"
)

type Eval interface {
	Eval(p *parser.Parser, ctx context.Context, t *base.T) (err error)
}

type MethodEvaluator struct {
	parser           *parser.Parser
	ctx              context.Context
	objectT          *base.T
	evaluatedObjectT *base.T
	method           string
	isParentheses    bool
	isAmpersand      bool
	outerEval        Eval
}

func NewMethodEvaluator(
	e Eval,
	p *parser.Parser,
	ctx context.Context,
	objectT *base.T,
	methodIdentifierT *base.T,
	isAmpersand bool,
) *MethodEvaluator {

	instance := objectT.ToString()
	var evaluatedObjectT *base.T

	switch instance {
	case "self":
		evaluatedObjectT = base.MakeObject(ctx.GetClass())

		if ctx.IsDefineStatic {
			objectT = base.MakeClass(ctx.GetClass())
		}

	default:
		evaluatedObjectT = objectT

		if objectT.IsIdentifierType() {
			evaluatedObjectT =
				base.GetDynamicValueT(
					ctx.GetFrame(),
					ctx.GetClass(),
					ctx.GetMethod(),
					instance,
				)
		}
	}

	p.SetLastEvaluatedT(evaluatedObjectT)

	p.SetLastCallFrameDetails(
		evaluatedObjectT.GetFrame(),
		evaluatedObjectT.GetObjectClass(),
		methodIdentifierT.ToString(),
	)

	key := evaluatedObjectT.GetFrame() + evaluatedObjectT.GetObjectClass() + methodIdentifierT.ToString()
	point := p.FileName + ":" + strconv.Itoa(p.Row)

	callPoint :=
		base.CallPoint{
			Point:        point,
			CallerFrame:  ctx.GetFrame(),
			CallerClass:  ctx.GetClass(),
			CallerMethod: ctx.GetMethod(),
		}

	if ctx.IsCheckRound() {
		base.MethodCallPoint[key] = append(base.MethodCallPoint[key], callPoint)
	}

	p.LastCallT = methodIdentifierT

	return &MethodEvaluator{
		outerEval:        e,
		parser:           p,
		ctx:              ctx,
		objectT:          objectT,
		evaluatedObjectT: evaluatedObjectT,
		method:           methodIdentifierT.ToString(),
		isAmpersand:      isAmpersand,
	}
}

func (m *MethodEvaluator) Evaluation() error {
	methodEvaluateStrategy := NewStrategy(m)

	return methodEvaluateStrategy.evaluate(m)
}
