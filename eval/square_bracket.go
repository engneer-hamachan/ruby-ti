package eval

import (
	"fmt"
	"ti/base"
	"ti/context"
	"ti/parser"
)

type SquareBracket struct{}

func NewSquareBracket() DynamicEvaluator {
	return &SquareBracket{}
}

func init() {
	DynamicEvaluators["["] = NewSquareBracket()
}

func isAcceptReferenceBind(ctx context.Context, objectT *base.T) bool {
	if objectT.IsUnionType() {
		return false
	}

	methodT :=
		base.GetMethodT(ctx.GetFrame(), ctx.GetClass(), objectT.ToString(), false)

	return methodT == nil
}

func (e *Evaluator) arrayReferenceEvaluation(
	p *parser.Parser,
	ctx context.Context,
	objectT *base.T,
	arrayT *base.T,
) error {

	for {
		id, isCloseParentheses, err := p.ReadWithCheck("]")
		if err != nil {
			return err
		}

		// a[1..]
		if id.IsTargetIdentifier("..") {
			base.SetValueT(
				ctx.GetFrame(),
				ctx.GetClass(),
				ctx.GetMethod(),
				objectT.ToString(),
				arrayT,
			)

			p.SetLastEvaluatedT(arrayT)
			p.Skip()
			return nil
		}

		if !isCloseParentheses {
			continue
		}

		break
	}

	_, isEquale, err := p.ReadWithCheck("=")
	if err != nil {
		return err
	}

	switch isEquale {
	// a[0] = 1
	case true:

		ctx.IsBind = true
		p.SkipNewline()

		nextT, err := p.Read()
		if err != nil {
			return err
		}

		methodT := base.GetMethodT(ctx.GetFrame(), "Array", "[]=", false)
		if methodT == nil {
			return fmt.Errorf("[]= is not defined method")
		}

		if !isAcceptReferenceBind(ctx, objectT) {
			return nil
		}

		err = e.Eval(p, ctx, nextT)
		if err != nil {
			return err
		}

		arrayT.AppendArrayVariant(p.GetLastEvaluatedT())

		base.SetValueT(
			ctx.GetFrame(),
			ctx.GetClass(),
			ctx.GetMethod(),
			objectT.ToString(),
			arrayT,
		)

		p.SetLastEvaluatedT(arrayT)

		return nil

	// a[0]
	default:
		methodT := base.GetMethodT(ctx.GetFrame(), "Array", "[]", false)
		if methodT == nil {
			return fmt.Errorf("[] is not defined method")
		}

		p.Unget()

		p.SetLastEvaluatedT(arrayT.UnifyVariants())

		return e.evalPriorityExp(p, ctx)
	}
}

func (e *Evaluator) hashReferenceEvaluation(
	p *parser.Parser,
	ctx context.Context,
	objectT *base.T,
	hashT *base.T,
) error {

	keyT, err := p.Read()
	if err != nil {
		return err
	}

	_, ok, err := p.ReadWithCheck("]")
	if err != nil {
		return err
	}

	if !ok {
		return fmt.Errorf("syntax error")
	}

	_, isEquale, err := p.ReadWithCheck("=")
	if err != nil {
		return err
	}

	switch isEquale {
	// a[:b] = 1
	case true:
		ctx.IsBind = true

		p.SkipNewline()

		nextT, err := p.Read()
		if err != nil {
			return err
		}

		methodT := base.GetMethodT(ctx.GetFrame(), "Hash", "[]=", false)
		if methodT == nil {
			return fmt.Errorf("[]= is not defined method")
		}

		if !isAcceptReferenceBind(ctx, objectT) {
			return nil
		}

		err = e.Eval(p, ctx, nextT)
		if err != nil {
			return err
		}

		valueT := p.GetLastEvaluatedT()

		var keyvalueT *base.T

		switch keyT.GetType() {
		case base.UNKNOWN:
			keyvalueT = base.MakeKeyValue(base.GenId(), &valueT)

		default:
			keyvalueT = base.MakeKeyValue(keyT.ToString(), &valueT)
		}

		hashT.AppendHashVariant(*keyvalueT)

		base.SetValueT(
			ctx.GetFrame(),
			ctx.GetClass(),
			ctx.GetMethod(),
			objectT.ToString(),
			hashT,
		)

		p.SetLastEvaluatedT(hashT)

		return nil

	// a[:b]
	default:
		methodT := base.GetMethodT(ctx.GetFrame(), "Hash", "[]", false)
		if methodT == nil {
			return fmt.Errorf("[] is not defined method")
		}

		p.Unget()

		switch keyT.GetType() {
		case base.UNKNOWN:
			p.SetLastEvaluatedT(hashT.UnifyVariants())
		default:
			p.SetLastEvaluatedT(hashT.HashReference(keyT.ToString()))
		}

		return e.evalPriorityExp(p, ctx)
	}
}

func (e *Evaluator) makeArray(
	p *parser.Parser,
	ctx context.Context,
	t *base.T,
) error {

	ctx.StartArrayCollect()
	defer ctx.EndArrayCollect()

	var isParentheses bool
	arrayT := base.MakeArray()

	if t.IsTargetIdentifier("[") {
		isParentheses = true
		p.SkipNewline()
	}

	if !isParentheses {
		err := e.Eval(p, ctx, t)
		if err != nil {
			return err
		}

		lastEvaluatedT := p.GetLastEvaluatedT()

		switch lastEvaluatedT.IsBeforeEvaluateAsteriskPrefix() {
		case true:
			for _, variant := range lastEvaluatedT.GetVariants() {
				arrayT.AppendArrayVariant(variant)
			}

		default:
			arrayT.AppendArrayVariant(lastEvaluatedT)
		}
	}

	for {
		nextT, err := p.Read()
		if err != nil {
			return err
		}

		if nextT == nil {
			break
		}

		if nextT.IsTargetIdentifier("]") && isParentheses {
			break
		}

		if nextT.IsTargetIdentifier("\n") && !isParentheses {
			p.Unget()
			break
		}

		if nextT.IsTargetIdentifier("\n") && isParentheses {
			continue
		}

		if nextT.IsTargetIdentifier(",") {
			p.SetLastEvaluatedT(base.MakeNil())
			continue
		}

		err = e.Eval(p, ctx, nextT)
		if err != nil {
			return err
		}

		if isParentheses {
			p.SkipNewline()
		}

		lastEvaluatedT := p.GetLastEvaluatedT()

		switch lastEvaluatedT.IsBeforeEvaluateAsteriskPrefix() {
		case true:
			for _, variant := range lastEvaluatedT.GetVariants() {
				arrayT.AppendArrayVariant(variant)
			}

		default:
			arrayT.AppendArrayVariant(lastEvaluatedT)
		}
	}

	p.SetLastEvaluatedT(arrayT)

	return nil
}

// a← current token pattern[]
func (e *Evaluator) referenceEvaluation(
	p *parser.Parser,
	ctx context.Context,
	objectT *base.T,
) error {

	objectIdentifier := objectT.ToString()

	t := base.GetDynamicValueT("", ctx.GetClass(), ctx.GetMethod(), objectIdentifier)

	if !t.IsArrayType() && !t.IsHashType() && !t.IsTargetClassObject("Proc") {
		p.SkipToTargetToken("]")

		return fmt.Errorf(
			"type mismatch. %s is not Array or Hash",
			objectT.ToString(),
		)
	}

	switch t.GetType() {
	case base.ARRAY:
		return e.arrayReferenceEvaluation(p, ctx, objectT, t)

	case base.HASH:
		return e.hashReferenceEvaluation(p, ctx, objectT, t)
	}

	p.SkipToTargetToken("]")
	p.SetLastEvaluatedT(base.MakeUntyped())

	return nil
}

// a[← current token pattern]
func (s *SquareBracket) Evaluation(
	e *Evaluator,
	p *parser.Parser,
	ctx context.Context,
	t *base.T,
) (err error) {

	lastT := p.GetLastEvaluatedT()

	// a[:b]
	if lastT.IsHashType() && p.IsParsingExpression() && !t.IsBeforeSpace {
		return e.hashReferenceEvaluation(p, ctx, base.MakeUnknown(), &lastT)
	}

	// a[0]
	if lastT.IsArrayType() && p.IsParsingExpression() && !t.IsBeforeSpace {
		return e.arrayReferenceEvaluation(p, ctx, base.MakeUnknown(), &lastT)
	}

	// []
	return e.makeArray(p, ctx, t)
}
