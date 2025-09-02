package eval

import (
	"ti/base"
	"ti/context"
	"ti/parser"
)

type Rescue struct{}

func NewRescue() DynamicEvaluator {
	return &Rescue{}
}

func init() {
	DynamicEvaluators["rescue"] = NewRescue()
}

func (r *Rescue) Evaluation(
	e *Evaluator,
	p *parser.Parser,
	ctx context.Context,
	t *base.T,
) (err error) {

	for {
		nextT, err := p.Read()
		if err != nil {
			return err
		}

		if nextT.IsNewLineIdentifier() {
			break
		}

		if nextT.IsTargetIdentifier("=>") {
			continue
		}

		base.SetValueT(
			ctx.GetFrame(),
			ctx.GetClass(),
			ctx.GetMethod(),
			nextT.ToString(),
			base.MakeObject("Bot"),
		)
	}

	return nil
}
