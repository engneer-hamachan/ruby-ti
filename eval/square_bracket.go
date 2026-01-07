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

	defineRow := p.ErrorRow

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
				ctx.IsDefineStatic,
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
		setDefineInfos(p, defineRow)

		base.SetValueT(
			ctx.GetFrame(),
			ctx.GetClass(),
			ctx.GetMethod(),
			objectT.ToString(),
			arrayT,
			ctx.IsDefineStatic,
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

	defineRow := p.ErrorRow

	nextT, err := p.Read()
	if err != nil {
		return err
	}

	// skip to end braquet
	for {
		err = e.Eval(p, ctx, nextT)
		if err != nil {
			return err
		}

		nextT, err := p.Read()
		if err != nil {
			return err
		}

		if nextT.IsTargetIdentifier("]") || nextT == nil {
			break
		}
	}

	keyT := p.GetLastEvaluatedT()

	_, isEquale, err := p.ReadWithCheck("=")
	if err != nil {
		return err
	}

	switch isEquale {
	// a[:b] = 1
	case true:
		ctx.IsBind = true
		p.EndParsingExpression()

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
		setDefineInfos(p, defineRow)

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
			ctx.IsDefineStatic,
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

func (e *Evaluator) stringReferenceEvaluation(
	p *parser.Parser,
	ctx context.Context,
	objectT *base.T,
	stringT *base.T,
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
				stringT,
				ctx.IsDefineStatic,
			)

			p.SetLastEvaluatedT(stringT)
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
		p.EndParsingExpression()

		methodT := base.GetMethodT(ctx.GetFrame(), "String", "[]=", false)
		if methodT == nil {
			return fmt.Errorf("[]= is not defined method")
		}

		ctx.IsBind = true
		p.SkipNewline()

		nextT, err := p.Read()
		if err != nil {
			return err
		}

		err = e.Eval(p, ctx, nextT)
		if err != nil {
			return err
		}

		base.SetValueT(
			ctx.GetFrame(),
			ctx.GetClass(),
			ctx.GetMethod(),
			objectT.ToString(),
			stringT,
			ctx.IsDefineStatic,
		)

		p.SetLastEvaluatedT(stringT)

		return nil

	// a[0]
	default:
		methodT := base.GetMethodT(ctx.GetFrame(), "String", "[]", false)
		if methodT == nil {
			return fmt.Errorf("[] is not defined method")
		}

		p.Unget()

		p.SetLastEvaluatedT(stringT)

		return e.evalPriorityExp(p, ctx)
	}
}

func (e *Evaluator) integerReferenceEvaluation(
	p *parser.Parser,
	ctx context.Context,
	objectT *base.T,
	intT *base.T,
) error {

	for {
		_, isCloseParentheses, err := p.ReadWithCheck("]")
		if err != nil {
			return err
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
		methodT := base.GetMethodT(ctx.GetFrame(), "Integer", "[]=", false)
		if methodT == nil {
			return fmt.Errorf("[]= is not defined method")
		}

		ctx.IsBind = true
		p.SkipNewline()

		nextT, err := p.Read()
		if err != nil {
			return err
		}

		err = e.Eval(p, ctx, nextT)
		if err != nil {
			return err
		}

		base.SetValueT(
			ctx.GetFrame(),
			ctx.GetClass(),
			ctx.GetMethod(),
			objectT.ToString(),
			intT,
			ctx.IsDefineStatic,
		)

		p.SetLastEvaluatedT(intT)

		return nil

	// a[0]
	default:
		methodT := base.GetMethodT(ctx.GetFrame(), "Integer", "[]", false)
		if methodT == nil {
			return fmt.Errorf("[] is not defined method")
		}

		p.Unget()

		p.SetLastEvaluatedT(intT)

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
	arrayT := base.MakeAnyArray()

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

	var t *base.T

	switch objectT.GetType() {
	case base.ARRAY, base.HASH, base.STRING:
		t = objectT

	case base.CONST:
		t = base.GetConstValueT(ctx.GetFrame(), ctx.GetClass(), objectT.ToString())

	default:
		objectIdentifier := objectT.ToString()

		t =
			base.GetDynamicValueT(
				"",
				ctx.GetClass(),
				ctx.GetMethod(),
				objectIdentifier,
			)
	}

	switch t.GetType() {
	case base.UNKNOWN, base.UNTYPED:
		p.SkipToTargetToken("]")
		p.SetLastEvaluatedT(base.MakeUntyped())

		return nil

	case base.ARRAY:
		return e.arrayReferenceEvaluation(p, ctx, objectT, t)

	case base.HASH:
		return e.hashReferenceEvaluation(p, ctx, objectT, t)

	case base.STRING:
		return e.stringReferenceEvaluation(p, ctx, objectT, t)

	case base.INT:
		return e.integerReferenceEvaluation(p, ctx, objectT, t)

	default:
		p.SkipToTargetToken("]")

		if t.IsTargetClassObject("Proc") {
			p.SetLastEvaluatedT(base.MakeUntyped())
			return nil
		}

		return fmt.Errorf(
			"type mismatch. %s is not Array or Hash",
			objectT.ToString(),
		)
	}
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

	if lastT.IsAnyType() && p.IsParsingExpression() && !t.IsBeforeSpace {
		p.SkipToTargetToken("]")
		p.SetLastEvaluatedT(base.MakeUntyped())
		return nil
	}

	// []
	return e.makeArray(p, ctx, t)
}
