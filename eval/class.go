package eval

import (
	"fmt"
	"slices"
	"ti/base"
	"ti/context"
	"ti/parser"
)

type Class struct{}

func NewClass() DynamicEvaluator {
	return &Class{}
}

func init() {
	DynamicEvaluators["class"] = NewClass()
}

func (c *Class) classIdentifierProcessing(
	e *Evaluator,
	p *parser.Parser,
	ctx *context.Context,
) error {

	_, b, err := p.ReadWithCheck("<<")
	if err != nil {
		p.Fatal(*ctx, err)
	}

	if !b {
		return nil
	}

	_, b, err = p.ReadWithCheck("self")
	if err != nil {
		p.Fatal(*ctx, err)
	}

	if !b {
		return fmt.Errorf("syntax error")
	}

	ctx.StartDefineStatic()
	defer ctx.EndDefineStatic()

	for {
		nextT, err := p.Read()
		if err != nil {
			p.Fatal(*ctx, err)
		}

		if nextT.IsEndIdentifier() {
			break
		}

		if nextT.IsTargetIdentifier("private") {
			ctx.StartPrivate()
			defer ctx.EndPrivate()

			continue
		}

		err = e.Eval(p, *ctx, nextT)
		if err != nil {
			p.Fatal(*ctx, err)
		}
	}

	return nil
}

func (c *Class) getNextFrame(ctx context.Context) string {
	if ctx.GetFrame() != "" && ctx.GetClass() != "" {
		return ctx.GetFrame() + "::" + ctx.GetClass()
	}

	return ctx.GetClass()
}

func (c *Class) Evaluation(
	e *Evaluator,
	p *parser.Parser,
	ctx context.Context,
	t *base.T,
) (err error) {

	nextT, err := p.Read()
	if err != nil {
		return err
	}

	nextFrame := c.getNextFrame(ctx)
	class := nextT.ToString()

	// make new method
	returnT := base.MakeObject(class)
	methodT := base.MakeMethod(nextFrame, "new", *returnT, []string{})
	base.SetClassMethodT(nextFrame, class, methodT, ctx.IsPrivate)

	// include ObjectClass
	if ctx.IsCollectRound() {
		classNode := base.ClassNode{Frame: ctx.GetFrame(), Class: class}
		objectClassNode := base.ClassNode{Frame: "Builtin", Class: ""}

		base.ClassInheritanceMap[classNode] =
			append(base.ClassInheritanceMap[classNode], objectClassNode)
	}

	ctx.SetFrame(nextFrame)
	ctx.SetClass(class)

	//extends
	nextT, err = p.Read()
	if err != nil {
		p.Fatal(ctx, err)
	}

	switch nextT.ToString() {
	case "<":
		nextT, err := p.Read()
		if err != nil {
			p.Fatal(ctx, err)
		}

		classNode := base.ClassNode{Frame: ctx.GetFrame(), Class: ctx.GetClass()}

		parentFrame, parentNamespace, parentClass := nextT.SeparateNameSpaces()

		parentFrame = base.CalculateFrame(parentFrame, parentNamespace)
		parentNode := base.ClassNode{Frame: parentFrame, Class: parentClass}

		if slices.Contains(base.ClassInheritanceMap[classNode], parentNode) {
			return nil
		}

		base.ClassInheritanceMap[classNode] =
			append(base.ClassInheritanceMap[classNode], parentNode)

	default:
		p.Unget()
	}

	for {
		nextT, err := p.Read()
		if err != nil {
			return err
		}

		if nextT.IsEndIdentifier() {
			break
		}

		if nextT.IsTargetIdentifier("class") {
			t, err := p.Read()
			if err != nil {
				return err
			}

			p.Unget()

			if t.IsTargetIdentifier("<<") {
				err = c.classIdentifierProcessing(e, p, &ctx)
				if err != nil {
					p.Fatal(ctx, err)
				}

				continue
			}
		}

		if nextT.IsTargetIdentifier("private") {
			methodT := base.GetMethodT("Builtin", "", "private", false)
			if methodT != nil {
				ctx.StartPrivate()
				defer ctx.EndPrivate()
			}

			continue
		}

		err = e.Eval(p, ctx, nextT)
		if err != nil {
			p.Fatal(ctx, err)
		}
	}

	return nil
}
