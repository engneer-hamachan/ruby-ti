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

func (d *Def) isComingArgs(t *base.T) bool {
	return t.IsOpenParentheses() || !t.IsNewLineIdentifier()
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

	if nextT.IsNewLineIdentifier() {
		if ctx.IsCollectRound() {
			base.SetValueT(
				ctx.GetFrame(),
				ctx.GetClass(),
				ctx.GetMethod(),
				argT.ToRemoveSuffixString(),
				base.MakeUnknown(),
				ctx.IsDefineStatic,
			)
		}

		argVariables = append(argVariables, argT.ToString())

		p.Unget()

		return argVariables, true, nil
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
				ctx.IsDefineStatic,
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
			ctx.IsDefineStatic,
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
	nextT *base.T,
) ([]string, bool, error) {

	var argVariables []string

	if !d.isComingArgs(nextT) {
		return argVariables, false, nil
	}

	isParentheses := nextT.IsOpenParentheses()

	if !isParentheses {
		p.Unget()
	}

	ctx.StartDefineArg()
	defer ctx.EndDefineArg()

	var asteriskCount int
	var doubleAsteriskCount int
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

		if doubleAsteriskCount > 0 && !argT.IsAmpersandPrefix() {
			p.Skip()
			return argVariables, false, fmt.Errorf(
				"cannot define parameter after '**' for %s",
				ctx.GetMethod(),
			)
		}

		if argT.IsDoubleAsteriskPrefix() {
			currentT :=
				base.GetValueT(
					ctx.GetFrame(),
					ctx.GetClass(),
					method,
					argT.ToString()[2:],
					ctx.IsDefineStatic,
				)

			if currentT == nil {
				base.SetValueT(
					ctx.GetFrame(),
					ctx.GetClass(),
					method,
					argT.ToString()[2:],
					base.MakeAnyHash(),
					ctx.IsDefineStatic,
				)
			}

			doubleAsteriskCount++
		}

		if argT.IsAsteriskPrefix() {
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
				ctx.IsDefineStatic,
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

func (d *Def) getMethodNameAndSetIsStatic(
	p *parser.Parser,
	ctx *context.Context,
) (string, error) {

	t, err := p.Read()
	if err != nil {
		return "", err
	}

	if t.IsTargetIdentifier("self") {
		ctx.IsDefineStatic = true

		t, err = p.ReadTwice()
		if err != nil {
			return "", err
		}
	}

	nextT, err := p.ReadAhead()
	if err != nil {
		return "", err
	}

	// def object.special_method
	if nextT.IsDotIdentifier() {
		objectT :=
			base.GetValueT(
				ctx.GetFrame(),
				ctx.GetClass(),
				ctx.GetMethod(),
				t.ToString(),
				ctx.IsDefineStatic,
			)

		if objectT.ID == "" {
			objectT.ID = base.GenId()
		}

		ctx.SetClass(objectT.ID)
		ctx.SetFrame(objectT.GetFrame())

		t, err = p.ReadTwice()
		if err != nil {
			return "", err
		}
	}

	method := t.ToString()

	if method == "initialize" {
		method = "new"
		ctx.IsDefineStatic = true
	}

	return method, nil
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
	defineRow int,
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

	methodT := d.makeDefineMethodT(p, ctx, method, args, returnT, false)

	d.setDefineMethodT(p, ctx, methodT, defineRow)

	return nil
}

func (d *Def) makeDefineMethodT(
	p *parser.Parser,
	ctx context.Context,
	method string,
	args []string,
	returnT base.T,
	isBlockGiven bool,
) *base.T {

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
	methodT.IsStatic = ctx.IsDefineStatic

	if isBlockGiven {
		methodT.IsBlockGiven = isBlockGiven
	}

	p.ClearTmpBlockParameters()

	return methodT
}

func (d *Def) setDefineMethodT(
	p *parser.Parser,
	ctx context.Context,
	methodT *base.T,
	defineRow int,
) {

	key := ctx.GetFrame() + ctx.GetClass() + methodT.GetMethodName()
	base.TSignatureDocument[key] = p.ConsumeTiComment()

	switch ctx.IsDefineStatic {
	case true:
		base.SetClassMethodT(
			ctx.GetFrame(),
			ctx.GetClass(),
			methodT,
			ctx.IsPrivate,
			p.FileName,
			defineRow,
		)

	default:
		// this proccess for not instance variable override check
		t :=
			base.GetInstanceValueT(
				ctx.GetFrame(),
				ctx.GetClass(),
				methodT.GetMethodName(),
			)

		if t.IsBeforeEvaluateAtmarkPrefix() && methodT.IsIdentifierType() {
			break
		}

		base.SetMethodT(
			ctx.GetFrame(),
			ctx.GetClass(),
			methodT,
			ctx.IsPrivate,
			p.FileName,
			defineRow,
		)
	}
}

func (d *Def) setDefineInfos(
	p *parser.Parser,
	ctx context.Context,
	methodT *base.T,
	defineRow int,
) {

	var prefix string

	prefix += "@"
	prefix += p.FileName + ":::"
	prefix += fmt.Sprintf("%d", defineRow)
	prefix += ":::"

	content :=
		base.MakeSignatureContent(prefix, ctx.GetFrame(), ctx.GetClass(), methodT)

	content += " ["

	switch ctx.IsDefineStatic {
	case true:
		content += "c/"
	default:
		content += "i/"
	}

	readable := [2]bool{ctx.IsPrivate, ctx.IsProtected}

	switch readable {
	case [2]bool{true, false}:
		content += "private"
	case [2]bool{false, true}:
		content += "protected"
	case [2]bool{false, false}:
		content += "public"
	}

	content += "]"

	p.DefineInfos = append(p.DefineInfos, content)
}

func (d *Def) makeReturnT(
	e *Evaluator,
	p *parser.Parser,
	ctx context.Context,
	method string,
) base.T {

	switch method {
	case "new":
		ctx.IsDefineStatic = true
		return *base.MakeObject(ctx.GetClass())

	default:
		returnT := d.getLastEvaluatedTWhenDefineMethod(e, p, ctx)
		returnT.SetBlockParamaters([]base.T{})
		returnT.IsBlockGiven = false

		return returnT
	}
}

func (d *Def) prepareParserSetting(p *parser.Parser, t *base.T) {
	p.ConsumeLastReturnT()
	p.EndParsingExpression()
	p.SetLastEvaluatedT(base.MakeNil())
	p.LastCallT = t
}

func (d *Def) Evaluation(
	e *Evaluator,
	p *parser.Parser,
	ctx context.Context,
	t *base.T,
) (err error) {

	defineRow := p.ErrorRow

	method, err := d.getMethodNameAndSetIsStatic(p, &ctx)
	if err != nil {
		p.Fatal(ctx, err)
	}

	ctx.SetMethod(method)
	d.prepareParserSetting(p, t)

	nextT, err := p.Read()
	if err != nil {
		p.Fatal(ctx, err)
	}

	// def hoge = 1
	if nextT.IsEqualIdentifier() && nextT.IsBeforeSpace {
		return d.endlessDefinition(e, p, ctx, method, []string{}, defineRow)
	}

	var args []string
	var isBlockGiven bool

	args, isBlockGiven, err =
		d.makeDefineArgVariables(e, p, ctx, method, nextT)

	if err != nil {
		p.Fatal(ctx, err)
	}

	nextT, err = p.Read()
	if err != nil {
		return err
	}

	// def hoge() = 1
	if nextT.IsEqualIdentifier() {
		return d.endlessDefinition(e, p, ctx, method, args, defineRow)
	}

	p.Unget()

	base.SnapShotArgumentTypes(ctx, method, args)

	err = d.evaluationBody(e, p, ctx)
	if err != nil && ctx.IsCheckRound() {
		p.Fatal(ctx, err)
	}

	base.RestoreArgumentTypes()

	returnT := d.makeReturnT(e, p, ctx, method)
	methodT := d.makeDefineMethodT(p, ctx, method, args, returnT, isBlockGiven)

	d.setDefineMethodT(p, ctx, methodT, defineRow)

	if ctx.IsCheckRound() {
		d.setDefineInfos(p, ctx, methodT, defineRow)
	}

	return nil
}
