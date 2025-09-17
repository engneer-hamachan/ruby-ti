package eval

import (
	"fmt"
	"ti/base"
	"ti/context"
	"ti/parser"
)

type Module struct{}

func NewModule() DynamicEvaluator {
	return &Module{}
}

func init() {
	DynamicEvaluators["module"] = NewModule()
}

func (m *Module) classIdentifierProcessing(
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

func (m *Module) getNextFrame(ctx context.Context) string {
	if ctx.GetFrame() != "" && ctx.GetClass() != "" {
		return ctx.GetFrame() + "::" + ctx.GetClass()
	}

	return ctx.GetClass()
}

func (m *Module) Evaluation(
	e *Evaluator,
	p *parser.Parser,
	ctx context.Context,
	t *base.T,
) (err error) {

	nextT, err := p.Read()
	if err != nil {
		p.Fatal(ctx, err)
	}

	nextFrame := m.getNextFrame(ctx)
	class := nextT.ToString()

	ctx.SetFrame(nextFrame)
	ctx.SetClass(class)

	for {
		nextT, err := p.Read()
		if err != nil {
			p.Fatal(ctx, err)
		}

		if nextT.IsEndIdentifier() {
			break
		}

		if nextT.IsTargetIdentifier("class") {
			t, err := p.Read()
			if err != nil {
				p.Fatal(ctx, err)
			}

			p.Unget()

			if t.IsTargetIdentifier("<<") {
				err = m.classIdentifierProcessing(e, p, &ctx)
				if err != nil {
					p.Fatal(ctx, err)
				}

				continue
			}
		}

		if nextT.IsTargetIdentifier("private") {
			ctx.StartPrivate()
			defer ctx.EndPrivate()

			continue
		}

		err = e.Eval(p, ctx, nextT)
		if err != nil {
			p.Fatal(ctx, err)
		}
	}

	// set defined class
	base.SetDefinedClass(nextFrame, class)

	return nil
}
