package eval

import (
	"fmt"
	"ti/base"
	"ti/context"
	"ti/parser"
)

type Def struct{}

func NewDef() DynamicEvaluator {
	return &Def{}
}

func init() {
	DynamicEvaluators["def"] = NewDef()
}

func (d *Def) bindDefaultKeywordArgs(
	e *Evaluator,
	p *parser.Parser,
	ctx context.Context,
	argT *base.T,
	argVariables []string,
) ([]string, bool, error) {

	nextT, err := p.Read()
	if err != nil {
		return []string{}, false, err
	}

	if !nextT.IsCloseParentheses() && !nextT.IsCommaIdentifier() {
		e.Eval(p, ctx, nextT)

		lastEvaluatedT := p.GetLastEvaluatedT()
		lastEvaluatedT.SetHasDefault(true)

		if ctx.IsCollectRound() {
			base.SetValueT(
				ctx.GetFrame(),
				ctx.GetClass(),
				ctx.GetMethod(),
				argT.ToRemoveSuffixString(),
				&lastEvaluatedT,
			)
		}

		argVariables = append(argVariables, argT.ToString())
		return argVariables, true, nil
	}

	p.Unget()

	return argVariables, false, nil
}

func (d *Def) bindDefaultArgs(
	e *Evaluator,
	p *parser.Parser,
	ctx context.Context,
	argVariables []string,
) ([]string, error) {

	leftT := p.GetLastEvaluatedT()

	nextT, err := p.Read()
	if err != nil {
		return argVariables, err
	}

	e.Eval(p, ctx, nextT)

	rightT := p.GetLastEvaluatedT()
	rightT.SetHasDefault(true)

	if ctx.IsCollectRound() {
		base.SetValueT(
			ctx.GetFrame(),
			ctx.GetClass(),
			ctx.GetMethod(),
			leftT.ToString(),
			&rightT,
		)

		argVariables = append(argVariables, leftT.ToString())

		return argVariables, nil
	}

	//	argVariables = append(argVariables, leftT.GetBeforeEvaluateCode())

	return argVariables, nil
}

func (d *Def) makeDefineArgVariables(
	e *Evaluator,
	p *parser.Parser,
	ctx context.Context,
	t *base.T,
) ([]string, error) {

	var argVariables []string

	isParentheses := t.IsOpenParentheses()

	if !isParentheses {
		p.Unget()
	}

	ctx.StartDefineArg()
	defer ctx.EndDefineArg()

	var asteriskCount int

	for {

		argT, err := p.Read()
		if err != nil {
			return argVariables, err
		}

		if argT.IsTargetIdentifier(",") {
			continue
		}

		if argT.IsTargetIdentifier(")") && isParentheses {
			break
		}

		if argT.IsNewLineIdentifier() && !isParentheses {
			p.Unget()
			break
		}

		if argT.ToString()[0] == '*' {
			asteriskCount++

			if asteriskCount > 1 {
				p.Skip()
				return argVariables, fmt.Errorf(
					"syntax errror. %s is define multiple '*'",
					ctx.GetMethod(),
				)
			}
		}

		// x=10
		if argT.IsEqualIdentifier() {
			argVariables, err = d.bindDefaultArgs(e, p, ctx, argVariables)
			if err != nil {
				return argVariables, err
			}

			continue
		}

		// x: 10
		if argT.IsKeyIdentifier() {
			var isBind bool

			argVariables, isBind, err =
				d.bindDefaultKeywordArgs(e, p, ctx, argT, argVariables)

			if err != nil {
				return argVariables, err
			}

			if isBind {
				continue
			}
		}

		e.Eval(p, ctx, argT)
		argVariables = append(argVariables, argT.ToString())
	}

	return argVariables, nil
}

func (d *Def) evaluationBody(
	e *Evaluator,
	p *parser.Parser,
	ctx context.Context,
) (err error) {

	for {
		nextT, err := p.Read()
		if err != nil {
			p.Fatal(ctx, err)
		}

		if nextT == nil {
			break
		}

		if nextT.IsEndIdentifier() {
			p.AppendLastReturnT()
			break
		}

		err = e.Eval(p, ctx, nextT)
		if err != nil {
			p.Fatal(ctx, err)
		}
	}

	return nil
}

func (d *Def) getMethodNameAndIsStatic(
	p *parser.Parser,
	ctx *context.Context,
) (string, bool, error) {

	var isStatic bool

	t, err := p.Read()
	if err != nil {
		return "", false, err
	}

	isStatic = ctx.IsDefineStatic

	if t.IsTargetIdentifier("self") {
		isStatic = true

		t, err = p.ReadTwice()
		if err != nil {
			return "", false, err
		}
	}

	// def hoge.fuga
	nextT, err := p.Read()
	if err != nil {
		return "", false, err
	}

	if nextT.IsDotIdentifier() {
		objectT :=
			base.GetValueT(
				ctx.GetFrame(),
				ctx.GetClass(),
				ctx.GetMethod(),
				t.ToString(),
			)

		if objectT.ID == "" {
			objectT.ID = base.GenId()
		}

		ctx.SetClass(objectT.ID)
		ctx.SetFrame(objectT.GetFrame())

		t, err = p.Read()
		if err != nil {
			return "", false, err
		}
	} else {
		p.Unget()
	}

	method := t.ToString()

	if method == "initialize" {
		method = "new"
	}

	return method, isStatic, nil
}

func (d *Def) getChainMethodReturnType(
	e *Evaluator,
	p *parser.Parser,
	ctx context.Context,
	evaluatedT base.T,
) base.T {

	for {
		if !evaluatedT.IsIdentifierType() {
			break
		}

		e.Eval(p, ctx, &evaluatedT)

		lastEvaluatedT := p.GetLastEvaluatedT()

		if lastEvaluatedT.IsTargetIdentifier(evaluatedT.ToString()) {
			break
		}

		evaluatedT = lastEvaluatedT

		continue
	}

	return evaluatedT
}

func (d *Def) unifyReturnT(lastReturnT []base.T) base.T {
	unionVariants := []base.T{}

	for _, t := range lastReturnT {
		if t.GetType() == base.UNKNOWN {
			continue
		}

		if t.GetType() != base.UNION {
			unionVariants = append(unionVariants, t)
			continue
		}

		unionVariants = append(unionVariants, t.GetVariants()...)
	}

	if len(unionVariants) == 1 {
		return unionVariants[0]
	}

	return *base.MakeUnifiedT(unionVariants)
}

func (d *Def) getLastEvaluatedTWhenDefineMethod(
	e *Evaluator,
	p *parser.Parser,
	ctx context.Context,
) base.T {

	var evaluatedT base.T
	lastReturnT := p.ConsumeLastReturnT()

	switch len(lastReturnT) {
	case 1:
		evaluatedT = lastReturnT[0]

	default:
		evaluatedT = d.unifyReturnT(lastReturnT)
	}

	return d.getChainMethodReturnType(e, p, ctx, evaluatedT)
}

func (d *Def) endlessDefiniction(
	e *Evaluator,
	p *parser.Parser,
	ctx context.Context,
	method string,
	args []string,
	isStatic bool,
) error {

	nextT, err := p.Read()
	if err != nil {
		return err
	}

	err = e.Eval(p, ctx, nextT)
	if err != nil {
		return err
	}

	returnT := p.GetLastEvaluatedT()

	methodT := base.MakeMethod(ctx.GetFrame(), method, returnT, args)
	methodT.SetBlockParamaters(p.GetTmpBlockParameters())

	p.ClearTmpBlockParameters()

	switch isStatic {
	case true:
		base.SetClassMethodT(ctx.GetFrame(), ctx.GetClass(), methodT, ctx.IsPrivate)

	default:
		base.SetMethodT(ctx.GetFrame(), ctx.GetClass(), methodT, ctx.IsPrivate)
	}

	return nil
}

func (d *Def) Evaluation(
	e *Evaluator,
	p *parser.Parser,
	ctx context.Context,
	t *base.T,
) (err error) {

	p.ConsumeLastReturnT()

	method, isStatic, err := d.getMethodNameAndIsStatic(p, &ctx)
	if err != nil {
		p.Fatal(ctx, err)
	}

	ctx.SetMethod(method)

	nextT, err := p.Read()
	if err != nil {
		p.Fatal(ctx, err)
	}

	if nextT.IsEqualIdentifier() {
		return d.endlessDefiniction(e, p, ctx, method, []string{}, isStatic)
	}

	var args []string

	if nextT.IsTargetIdentifier("(") || !nextT.IsTargetIdentifier("\n") {
		args, err = d.makeDefineArgVariables(e, p, ctx, nextT)
		if err != nil {
			p.Fatal(ctx, err)
		}
	}

	tmpArgs := []string{}
	isBlockGiven := false

	for _, arg := range args {
		if len(arg) > 1 && arg[0] == '&' {
			base.SetValueT(
				ctx.GetFrame(),
				ctx.GetClass(),
				method,
				arg[1:],
				base.MakeObject("Proc"),
			)

			isBlockGiven = true

			continue
		}

		tmpArgs = append(tmpArgs, arg)
	}

	args = tmpArgs

	nextT, err = p.Read()
	if err != nil {
		return err
	}

	if nextT.IsEqualIdentifier() {
		return d.endlessDefiniction(e, p, ctx, method, args, isStatic)
	}

	p.Unget()

	err = d.evaluationBody(e, p, ctx)
	if err != nil && ctx.IsCheckRound() {
		p.Fatal(ctx, err)
	}

	var returnT base.T

	switch method {
	case "new":
		isStatic = true
		returnT = *base.MakeObject(ctx.GetClass())

	default:
		returnT = d.getLastEvaluatedTWhenDefineMethod(e, p, ctx)
	}

	methodT := base.MakeMethod(ctx.GetFrame(), method, returnT, args)
	methodT.SetBlockParamaters(p.GetTmpBlockParameters())

	if isBlockGiven {
		methodT.IsBlockGiven = isBlockGiven
		methodT.SetBlockParamaters([]base.T{*base.MakeUntyped()})
	}

	p.ClearTmpBlockParameters()

	switch isStatic {
	case true:
		base.SetClassMethodT(ctx.GetFrame(), ctx.GetClass(), methodT, ctx.IsPrivate)

	default:
		base.SetMethodT(ctx.GetFrame(), ctx.GetClass(), methodT, ctx.IsPrivate)
	}

	return nil
}
