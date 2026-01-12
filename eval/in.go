package eval

import (
	"ti/base"
	"ti/context"
	"ti/parser"
)

type In struct{}

func NewIn() DynamicEvaluator {
	return &In{}
}

func init() {
	DynamicEvaluators["in"] = NewIn()
}

func (i *In) parseVariable(
	p *parser.Parser,
	ctx context.Context,
	t *base.T,
) error {

	var variable string
	var err error

	if t.IsKeyIdentifier() {
		variable = t.ToString()[:len(t.ToString())-1]
	} else {
		variable = t.ToString()
	}

	base.SetValueT(
		ctx.GetFrame(),
		ctx.GetClass(),
		ctx.GetMethod(),
		variable,
		base.MakeUntyped(),
		ctx.IsDefineStatic,
	)

	t, err = p.Read()
	if err != nil {
		return err
	}

	if t.IsCommaIdentifier() {
		t, err = p.Read()
		if err != nil {
			return nil
		}

		err = i.parsePattern(p, ctx, t)
		if err != nil {
			return err
		}

		return nil
	}

	p.Unget()

	return nil
}

func (i *In) parseHash(
	p *parser.Parser,
	ctx context.Context,
) error {

	for {
		nextT, err := p.Read()
		if err != nil {
			return err
		}

		if nextT.IsVariableIdentifier() {
			var variable string

			if nextT.IsKeyIdentifier() {
				variable = nextT.ToString()[:len(nextT.ToString())-1]
			} else {
				variable = nextT.ToString()
			}

			base.SetValueT(
				ctx.GetFrame(),
				ctx.GetClass(),
				ctx.GetMethod(),
				variable,
				base.MakeUntyped(),
				ctx.IsDefineStatic,
			)
		}

		if nextT.IsTargetIdentifier("}") {
			break
		}
	}

	return nil
}

func (i *In) parseArray(
	p *parser.Parser,
	ctx context.Context,
) error {

	for {
		nextT, err := p.Read()
		if err != nil {
			return err
		}

		if nextT.IsVariableIdentifier() {
			err = i.parsePattern(p, ctx, nextT)
			if err != nil {
				return err
			}
		}

		if nextT.IsTargetIdentifier("]") {
			break
		}
	}

	return nil
}

func (i *In) parseClass(
	p *parser.Parser,
	ctx context.Context,
	t *base.T,
) error {

	objectT := base.MakeObject(t.ToString())

	// =>
	nextT, err := p.Read()
	if err != nil {
		return err
	}

	if nextT.IsTargetIdentifier("=>") {
		// x
		nextT, err := p.Read()
		if err != nil {
			return err
		}

		base.SetValueT(
			ctx.GetFrame(),
			ctx.GetClass(),
			ctx.GetMethod(),
			nextT.ToString(),
			objectT,
			ctx.IsDefineStatic,
		)
	}

	return nil
}

func (i *In) parsePattern(
	p *parser.Parser,
	ctx context.Context,
	nextT *base.T,
) error {

	switch {
	// String => x
	case nextT.IsClassType():
		err := i.parseClass(p, ctx, nextT)
		if err != nil {
			return err
		}

	// [x, y]
	case nextT.IsTargetIdentifier("["):
		err := i.parseArray(p, ctx)
		if err != nil {
			return err
		}

	// {name:, age:}
	case nextT.IsTargetIdentifier("{"):
		err := i.parseHash(p, ctx)
		if err != nil {
			return err
		}

	// x
	case nextT.IsVariableIdentifier():
		err := i.parseVariable(p, ctx, nextT)
		if err != nil {
			return err
		}
	}

	return nil
}

func (i *In) Evaluation(
	e *Evaluator,
	p *parser.Parser,
	ctx context.Context,
	t *base.T,
) (err error) {

	// TODO:
	// 1. t.IsVariableIdentifier -> done
	// 2. implement bnf
	//   pattern :=
	//     literal
	//   | variable
	//   | _
	//   | [pattern, ...]
	//
	//   | { key: pattern, **pattern }
	//   | Class
	//   | Class[pattern, ...]
	//   | Range
	//   | pattern | pattern
	//   | pattern & pattern
	//   | ^variable
	//   | pattern if expr
	//   | pattern => variable
	// 3. array inference
	// 4. hash inference
	nextT, err := p.Read()
	if err != nil {
		return err
	}

	i.parsePattern(p, ctx, nextT)

	p.SkipToTargetToken("\n")

	return nil
}
