package eval

import (
	"ti/base"
	"ti/context"
	"ti/parser"
)

type While struct{}

func NewWhile() DynamicEvaluator {
	return &While{}
}

func init() {
	while := NewWhile()
	DynamicEvaluators["while"] = while
	DynamicEvaluators["until"] = while
	DynamicEvaluators["for"] = while
}

func (d *While) Evaluation(
	e *Evaluator,
	p *parser.Parser,
	ctx context.Context,
	t *base.T,
) (err error) {

	isEatedNewlineToken := false

	for {
		nextT, err := p.Read()
		if err != nil {
			return err
		}

		if nextT == nil {
			break
		}

		if nextT.IsNewLineIdentifier() {
			isEatedNewlineToken = true
		}

		if nextT.IsTargetIdentifier("do") && !isEatedNewlineToken {
			continue
		}

		if nextT.IsTargetIdentifier("{") && !isEatedNewlineToken {
			continue
		}

		if nextT.IsTargetIdentifier("end") || nextT.IsTargetIdentifier("}") {
			break
		}

		err = e.Eval(p, ctx, nextT)
		if err != nil && !ctx.IsCollectRound() {
			return err
		}
	}

	return nil
}
