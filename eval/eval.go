package eval

import (
	"fmt"
	"ti/base"
	"ti/context"
	"ti/parser"
)

type Evaluator struct{}

func (e *Evaluator) evalPriorityExp(
	p *parser.Parser,
	ctx context.Context,
) error {

	nextT, err := p.Read()
	if err != nil {
		return err
	}

	if nextT.IsPriorityT() {
		err = e.Eval(p, ctx, nextT)
		if err != nil {
			return err
		}

		return nil
	}

	p.Unget()

	return nil
}

func (e *Evaluator) setLastEvaluatedT(
	p *parser.Parser,
	ctx context.Context,
	t *base.T,
) {

	switch t.GetType() {
	case base.UNKNOWN:
		switch {
		case t.GetBeforeEvaluateCode() != "":
			p.SetLastEvaluatedT(t)

		case ctx.GetClass() != "" && ctx.GetMethod() != "":
			t.SetBeforeEvaluateCode(ctx.GetClass() + "." + ctx.GetMethod())
			p.SetLastEvaluatedT(t)

			return

		case ctx.GetClass() != "":
			t.SetBeforeEvaluateCode(ctx.GetMethod())
			p.SetLastEvaluatedT(t)

			return
		}

	case base.INT:
		if t.GetBeforeEvaluateCode() != "" {
			break
		}

		fmt.Println(t.ToString())
		t.SetBeforeEvaluateCode(t.ToString())

	case base.STRING:
		if t.GetBeforeEvaluateCode() != "" {
			break
		}

		t.SetBeforeEvaluateCode("'" + t.ToString() + "'")

	case base.NIL:
		if t.GetBeforeEvaluateCode() != "" {
			break
		}

		t.SetBeforeEvaluateCode("nil")

	default:
		if t.GetBeforeEvaluateCode() != "" {
			break
		}

		t.SetBeforeEvaluateCode(t.ToString())
	}

	p.SetLastEvaluatedT(t)
}

func (e *Evaluator) ContinuousEval(
	p *parser.Parser,
	ctx context.Context,
	t *base.T,
	continueIdentifiier string,
) error {

	err := e.Eval(p, ctx, t)
	if err != nil {
		return err
	}

	nextT, err := p.Read()
	if err != nil {
		return err
	}

	p.Unget()

	if nextT.IsTargetIdentifier(continueIdentifiier) {
		err = e.Eval(p, ctx, nextT)
		if err != nil {
			return err
		}
	}

	return nil
}

func (e *Evaluator) Eval(
	p *parser.Parser,
	ctx context.Context,
	t *base.T,
) (err error) {

	if t == nil {
		return nil
	}

	nextT, err := p.Read()
	if err != nil {
		return err
	}

	switch {
	//CONST
	case t.IsConstType():
		return e.handleConstEvaluation(p, ctx, t)

	// A::B
	case t.IsNameSpaceIdentifier():
		return nameSpaceEvaluation(e, p, ctx, t)

	// hoge[:a]
	case t.IsRefferenceAbleT() && nextT.IsRefferenceSquareT():
		return e.referenceEvaluation(p, ctx, t)

	// hoge.fuga
	case t.IsDotIdentifier():
		t := p.GetLastEvaluatedT()

		return e.handleEvaluateMethod(p, ctx, &t, nextT, false)

	// hoge.fuga
	case nextT.IsDotIdentifier() && !t.IsNewLineIdentifier():
		nextT, err := p.Read()
		if err != nil {
			return err
		}

		return e.handleEvaluateMethod(p, ctx, t, nextT, false)

	// hoge&.fuga
	case t.IsAndDotIdentifier():
		t := p.GetLastEvaluatedT()

		return e.handleEvaluateMethod(p, ctx, &t, nextT, true)

	// hoge&.fuga
	case nextT.IsAndDotIdentifier():
		nextT, err := p.Read()
		if err != nil {
			return err
		}

		return e.handleEvaluateMethod(p, ctx, t, nextT, true)

	// 1..2
	case nextT.IsTargetIdentifier("..") || nextT.IsTargetIdentifier("..."):
		err := e.Eval(p, ctx, t)
		if err != nil {
			return err
		}

		evaluator := DynamicEvaluators[nextT.ToString()]

		return evaluator.Evaluation(e, p, ctx, nextT)

	// { |
	case t.IsTargetIdentifier("{") && nextT.IsTargetIdentifier("|"):
		p.Unget()

	// do |
	case t.IsTargetIdentifier("do") && nextT.IsTargetIdentifier("|"):
		p.Unget()

	// 1 + 1
	case t.IsTransformTargetIdentifier():
		p.Unget()
		objectT := p.GetLastEvaluatedT()
		return e.handleEvaluateMethod(p, ctx, &objectT, t, false)

	// 1 + 1
	case p.LastCallT.IsNotPowerDown(nextT) && nextT.IsTransformTargetIdentifier() && !t.IsTargetIdentifier("def"):
		err := e.Eval(p, ctx, t)
		if err != nil {
			return err
		}

		objectT := p.GetLastEvaluatedT()
		return e.handleEvaluateMethod(p, ctx, &objectT, nextT, false)

	// 1
	case t.IsImmediate():
		p.Unget()
		e.setLastEvaluatedT(p, ctx, t)

		return nil

	// \n
	case t.IsNewLineIdentifier():
		p.Unget()
		p.EndParsingExpression()

		return nil

	//=begin
	case t.IsEqualIdentifier() && nextT.IsTargetIdentifier("begin"):
		return skipMultilineComment(p)

	// test()
	case t.IsTopLevelFunctionIdentifier(ctx.GetFrame(), ctx.GetClass()):
		p.Unget()
		return e.handleEvaluateMethod(p, ctx, base.MakeObjectObject(), t, false)

	default:
		p.Unget()
	}

	err = e.handleDynamicOrIdentifierEvaluator(p, ctx, t)
	if err != nil {
		return err
	}

	return nil
}
