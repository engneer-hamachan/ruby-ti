package eval

import (
	"ti/base"
	"ti/context"
	"ti/parser"
)

type NameSpace struct{}

func nameSpaceEvaluation(
	e *Evaluator,
	p *parser.Parser,
	ctx context.Context,
	t *base.T,
) (err error) {

	p.Unget()

	frame, parentClass, class := base.SeparateNameSpaces(t.ToString())

	t = base.MakeIdentifier(class)

	switch {
	case t.IsClassIdentifier():
		ctx.SetFrame(base.CalculateFrame(frame, parentClass))
		t = base.MakeClass(t.ToString())

	case t.IsConstIdentifier():
		ctx.SetFrame(base.CalculateFrame(frame, parentClass))
		t = base.MakeConst(t.ToString())

	default:
		ctx.SetFrame(frame)
		ctx.SetClass(parentClass)
	}

	err = e.Eval(p, ctx, t)
	if err != nil {
		return err
	}

	return nil
}
