package eval

import (
	"ti/base"
	"ti/context"
	"ti/parser"
)

type Hash struct{}

func NewHash() DynamicEvaluator {
	return &Hash{}
}

func init() {
	DynamicEvaluators["{"] = NewHash()
}

func (h *Hash) Evaluation(
	e *Evaluator,
	p *parser.Parser,
	ctx context.Context,
	t *base.T,
) (err error) {

	newHash := base.MakeAnyHash()

	for {
		p.SkipNewline()

		nextT, err := p.Read()
		if err != nil {
			return err
		}

		if nextT == nil || nextT.IsTargetIdentifier("}") {
			break
		}

		if nextT.IsCommaIdentifier() {
			continue
		}

		if nextT.IsTargetIdentifier("|") {
			continue
		}

		var key string

		switch nextT.GetType() {
		case base.UNKNOWN:
			key = ":" + nextT.ToString()[:len(nextT.ToString())-1]

		default:
			key = nextT.ToString()
		}

		nextT, err = p.Read()
		if err != nil {
			return err
		}

		if nextT.IsTargetIdentifier("=>") {
			nextT, err = p.Read()
			if err != nil {
				return err
			}
		}

		err = e.ContinuousEval(p, ctx, nextT, ".")
		if err != nil {
			return err
		}

		valueT := p.GetLastEvaluatedT()

		newHash.AppendHashVariant(*base.MakeKeyValue(key, &valueT))
	}

	p.SetLastEvaluatedT(newHash)

	return nil
}
