package eval

import (
	"ti/base"
	"ti/context"
	"ti/parser"
)

type Do struct{}

func NewDo() DynamicEvaluator {
	return &Do{}
}

func init() {
	DynamicEvaluators["do"] = NewDo()
}

func (d *Do) makeRestoreFunc(
	ctx context.Context,
	tFrame map[base.FrameKey]*base.T,
	restoreVariables []RestoreVariable,
) func() {

	return func() {
		base.RestoreFrame(base.TFrame, tFrame)

		for _, r := range restoreVariables {
			base.SetValueT(
				ctx.GetFrame(),
				ctx.GetClass(),
				ctx.GetMethod(),
				r.id,
				r.t,
			)
		}
	}
}

func (d *Do) setBlockParameters(
	p *parser.Parser,
	ctx context.Context,
	blockVariables []*base.T,
) {

	lastEvaluatedT := p.GetLastEvaluatedT()
	frame, class, method := p.GetLastCallFrameDetails()

	methodT := base.GetMethodT(frame, class, method, false)

	if methodT != nil {
		var blockParamaters []base.T

		for _, t := range methodT.GetBlockParameters() {
			if t.GetType() == base.UNIFY {
				blockParamaters =
					append(blockParamaters, *lastEvaluatedT.UnifyVariants())

				continue
			}

			if t.GetType() == base.SELF {
				blockParamaters = append(blockParamaters, lastEvaluatedT)
				continue
			}

			if t.GetType() == base.UNIFIED_SELF_ARGUMENT {
				tmpArgTs := p.GetTmpEvaluaetdArgs()
				blockParamaters = append(blockParamaters, *tmpArgTs[0].UnifyVariants())
				continue
			}

			blockParamaters = append(blockParamaters, t)
		}

		lastEvaluatedT.SetBlockParamaters(blockParamaters)
	}

	blockParameters := lastEvaluatedT.GetBlockParameters()

	if len(lastEvaluatedT.GetBlockParameters()) < 1 {
		return
	}

	for idx, variable := range blockVariables {
		if len(blockParameters) <= idx {
			base.SetValueT(
				ctx.GetFrame(),
				ctx.GetClass(),
				ctx.GetMethod(),
				variable.ToString(),
				base.MakeNil(),
			)

			continue
		}

		base.SetValueT(
			ctx.GetFrame(),
			ctx.GetClass(),
			ctx.GetMethod(),
			variable.ToString(),
			&blockParameters[idx],
		)
	}
}

func (d *Do) collectBlockVariables(p *parser.Parser) ([]*base.T, error) {
	nextT, ok, err := p.ReadWithCheck("|")
	if err != nil {
		return []*base.T{}, err
	}

	if !ok {
		if nextT.IsNewLineIdentifier() {
			p.Unget()
		}

		return []*base.T{}, err
	}

	var blockVariables []*base.T

	for {
		nextT, err := p.Read()
		if err != nil {
			return []*base.T{}, err
		}

		if nextT.IsTargetIdentifier("|") {
			return blockVariables, nil
		}

		if nextT.IsCommaIdentifier() {
			continue
		}

		blockVariables = append(blockVariables, nextT)
	}
}

type RestoreVariable struct {
	id string
	t  *base.T
}

func (d *Do) prepare(p *parser.Parser, ctx context.Context) (func(), error) {
	tFrame := base.DeepCopyTFrame()

	blockVariables, err := d.collectBlockVariables(p)
	if err != nil {
		return func() {}, err
	}

	var restoreVariables []RestoreVariable

	for _, id := range blockVariables {
		t :=
			base.GetValueT(
				ctx.GetFrame(),
				ctx.GetClass(),
				ctx.GetMethod(),
				id.ToString(),
			)

		restoreVariables =
			append(restoreVariables, RestoreVariable{id.ToString(), t})
	}

	d.setBlockParameters(p, ctx, blockVariables)

	return d.makeRestoreFunc(ctx, tFrame, restoreVariables), nil
}

func (d *Do) Evaluation(
	e *Evaluator,
	p *parser.Parser,
	ctx context.Context,
	t *base.T,
) (err error) {

	nextT, err := p.Read()
	if err != nil {
		return err
	}

	p.Unget()

	// {}
	if nextT.IsTargetIdentifier("}") {
		p.SetLastEvaluatedT(base.MakeBlock())
		return nil
	}

	restoreFunc, err := d.prepare(p, ctx)
	if err != nil {
		return err
	}

	defer restoreFunc()

	for {
		nextT, err := p.Read()
		if err != nil {
			return err
		}

		if nextT.IsTargetIdentifier("end") || nextT.IsTargetIdentifier("}") {
			break
		}

		err = e.Eval(p, ctx, nextT)
		if err != nil {
			return err
		}
	}

	lastEvaluatedT := p.GetLastEvaluatedT()
	p.SetLastEvaluatedT(base.MakeBlockWithResult(&lastEvaluatedT))

	return nil
}
