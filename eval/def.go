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
		return argVariables, false, err
	}

	if !nextT.IsCloseParentheses() && !nextT.IsCommaIdentifier() {
		err := e.Eval(p, ctx, nextT)
		if err != nil {
			return argVariables, false, err
		}

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
) error {

	leftT := p.GetLastEvaluatedT()

	nextT, err := p.Read()
	if err != nil {
		return err
	}

	err = e.Eval(p, ctx, nextT)
	if err != nil {
		return err
	}

	if nextT == nil {
		return nil
	}

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

		return nil
	}

	return nil
}

func (d *Def) makeDefineArgVariables(
	e *Evaluator,
	p *parser.Parser,
	ctx context.Context,
	method string,
	isParentheses bool,
) ([]string, bool, error) {

	var argVariables []string

	if !isParentheses {
		p.Unget()
	}

	ctx.StartDefineArg()
	defer ctx.EndDefineArg()

	var asteriskCount int
	isBlockGiven := false

	for {
		argT, err := p.Read()
		if err != nil {
			return argVariables, false, err
		}

		if argT.IsTargetIdentifier("end") {
			p.Unget()
			return argVariables, false, err
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
				return argVariables, false, fmt.Errorf(
					"syntax errror. %s is define multiple '*'",
					ctx.GetMethod(),
				)
			}
		}

		// x=10
		if argT.IsEqualIdentifier() {
			err = d.bindDefaultArgs(e, p, ctx)
			if err != nil {
				return argVariables, false, err
			}

			continue
		}

		// x: 10
		if argT.IsKeyIdentifier() {
			var isBind bool

			argVariables, isBind, err =
				d.bindDefaultKeywordArgs(e, p, ctx, argT, argVariables)

			if err != nil {
				return argVariables, false, err
			}

			if isBind {
				continue
			}
		}

		err = e.Eval(p, ctx, argT)
		if err != nil {
			return argVariables, false, err
		}

		arg := argT.ToString()

		// &block
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

		argVariables = append(argVariables, arg)
	}

	return argVariables, isBlockGiven, nil
}

func (d *Def) evaluationBody(
	e *Evaluator,
	p *parser.Parser,
	ctx context.Context,
) (err error) {

	p.SetLastEvaluatedT(base.MakeNil())
	p.EndParsingExpression()

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

	nextT, err := p.ReadAhead()
	if err != nil {
		return "", false, err
	}

	// def object.special_method
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

		t, err = p.ReadTwice()
		if err != nil {
			return "", false, err
		}
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

		if t.GetType() == base.UNION {
			unionVariants = append(unionVariants, t.GetVariants()...)
			continue
		}

		unionVariants = append(unionVariants, t)
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

func (d *Def) endlessDefinition(
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

func isKeySuffix(str string) bool {
	return str[len(str)-1:] == ":" && len(str) >= 2
}

func isAsteriskPrefix(str string) bool {
	return str[0] == '*'
}

func removeSuffix(str string) string {
	return str[:len(str)-1]
}

func (d *Def) setDefineInfos(
	p *parser.Parser,
	ctx context.Context,
	methodT *base.T,
	isStatic bool,
	defineRow int,
) {

	var hint string

	hint += "@"
	hint += p.FileName + ":::"
	hint += fmt.Sprintf("%d", defineRow)
	hint += ":::"

	argumentTypes := "("

	for _, definedArg := range methodT.GetDefineArgs() {
		if argumentTypes != "(" {
			argumentTypes += ", "
		}

		if isKeySuffix(definedArg) {
			argumentTypes += definedArg + " "
			definedArg = removeSuffix(definedArg)
		}

		if isAsteriskPrefix(definedArg) {
			argumentTypes += "*"
		}

		definedArgT :=
			base.GetValueT(
				methodT.GetFrame(),
				ctx.GetClass(),
				methodT.GetMethodName(),
				definedArg,
			)

		if definedArgT.HasDefault() && !definedArgT.IsUnionType() {
			argumentTypes += "?"
		}

		if isAsteriskPrefix(definedArg) {
			definedArgT = definedArgT.UnifyVariants()
		}

		switch definedArgT.GetType() {
		case base.UNION:
			argumentTypes += base.UnionTypeToString(definedArgT.GetVariants())

		case base.UNKNOWN:
			argumentTypes += "?"

		default:
			argumentTypes += base.TypeToString(definedArgT)
		}
	}

	argumentTypes += ")"

	hint += argumentTypes

	if methodT.IsBlockGiven {
		hint += " <is_block_given: true>"
	}

	hint += " -> "

	switch methodT.GetType() {
	case base.UNION:
		hint += base.UnionTypeToString(methodT.GetVariants())

	case base.UNKNOWN:
		hint += "?"

	default:
		hint += base.TypeToString(methodT)
	}

	hint += " ["

	switch isStatic {
	case true:
		hint += "c/"
	default:
		hint += "i/"
	}

	readable := [2]bool{ctx.IsPrivate, ctx.IsProtected}

	switch readable {
	case [2]bool{true, false}:
		hint += "private"
	case [2]bool{false, true}:
		hint += "protected"
	case [2]bool{false, false}:
		hint += "public"
	}

	hint += "]"

	p.DefineInfos = append(p.DefineInfos, hint)
}

func (d *Def) Evaluation(
	e *Evaluator,
	p *parser.Parser,
	ctx context.Context,
	t *base.T,
) (err error) {

	p.LastCallT = t
	p.ConsumeLastReturnT()
	p.SetLastEvaluatedT(base.MakeNil())
	p.EndParsingExpression()
	defineRow := p.ErrorRow

	method, isStatic, err := d.getMethodNameAndIsStatic(p, &ctx)
	if err != nil {
		p.Fatal(ctx, err)
	}

	ctx.SetMethod(method)

	nextT, err := p.Read()
	if err != nil {
		p.Fatal(ctx, err)
	}

	// def hoge = 1
	if nextT.IsEqualIdentifier() && nextT.IsBeforeSpace {
		return d.endlessDefinition(e, p, ctx, method, []string{}, isStatic)
	}

	var args []string
	var isBlockGiven bool

	if nextT.IsTargetIdentifier("(") || !nextT.IsTargetIdentifier("\n") {
		args, isBlockGiven, err =
			d.makeDefineArgVariables(e, p, ctx, method, nextT.IsOpenParentheses())

		if err != nil {
			p.Fatal(ctx, err)
		}
	}

	nextT, err = p.Read()
	if err != nil {
		return err
	}

	// def hoge() = 1
	if nextT.IsEqualIdentifier() {
		return d.endlessDefinition(e, p, ctx, method, args, isStatic)
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

	var methodT *base.T

	switch returnT.GetType() {
	case base.OBJECT:
		methodT = base.MakeMethod(returnT.GetFrame(), method, returnT, args)

	default:
		methodT = base.MakeMethod(ctx.GetFrame(), method, returnT, args)
	}

	methodT.SetBlockParamaters(p.GetTmpBlockParameters())
	methodT.IsProtected = ctx.IsProtected
	methodT.DefinedFrame = ctx.GetFrame()
	methodT.DefinedClass = ctx.GetClass()

	if isBlockGiven {
		methodT.IsBlockGiven = isBlockGiven
		methodT.SetBlockParamaters([]base.T{*base.MakeUntyped()})
	}

	p.ClearTmpBlockParameters()

	switch isStatic {
	case true:
		base.SetClassMethodT(ctx.GetFrame(), ctx.GetClass(), methodT, ctx.IsPrivate)

	default:
		// this proccess for not instance variable override check
		t = base.GetInstanceValueT(ctx.GetFrame(), ctx.GetClass(), method)
		if t != nil && t.IsBeforeEvaluateAtmarkPrefix() && methodT.IsIdentifierType() {
			break
		}

		base.SetMethodT(ctx.GetFrame(), ctx.GetClass(), methodT, ctx.IsPrivate)
	}

	if ctx.IsCheckRound() {
		d.setDefineInfos(p, ctx, methodT, isStatic, defineRow)
	}

	return nil
}
