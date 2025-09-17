package eval

import (
	"ti/base"
	"ti/context"
	"ti/parser"
)

type Comma struct{}

func NewComma() DynamicEvaluator {
	return &Comma{}
}

func init() {
	DynamicEvaluators[","] = NewComma()
}

func (d *Comma) Evaluation(
	e *Evaluator,
	p *parser.Parser,
	ctx context.Context,
	t *base.T,
) (err error) {

	var tArray []*base.T
	tArray = append(tArray, p.GetLastEvaluatedTPointer().(*base.T))

	for {
		nextT, err := p.Read()
		if err != nil {
			return err
		}

		if nextT == nil || nextT.IsNewLineIdentifier() {
			p.Unget()
			break
		}

		if nextT.IsCommaIdentifier() {
			continue
		}

		if nextT.IsEqualIdentifier() {
			p.Unget()
			p.SetLastEvaluatedT(tArray)

			return nil
		}

		err = e.Eval(p, ctx, nextT)
		if err != nil {
			return err
		}

		tArray = append(tArray, p.GetLastEvaluatedTPointer().(*base.T))

		nextT, err = p.Read()
		if err != nil {
			return err
		}

		p.Unget()

		if nextT.IsCommaIdentifier() || nextT.IsEqualIdentifier() {
			continue
		}

		break
	}

	nextT, err := p.Read()
	if err != nil {
		return err
	}

	p.Unget()

	if nextT.IsEqualIdentifier() {
		p.SetLastEvaluatedT(tArray)

		return nil
	}

	arrayT := base.MakeArray()

	for _, t := range tArray {
		if t == nil {
			arrayT.AppendArrayVariant(*base.MakeNil())
			continue
		}

		if t.IsBeforeEvaluateAsteriskPrefix() {
			for _, variant := range t.GetVariants() {
				arrayT.AppendArrayVariant(variant)
			}

			continue
		}

		arrayT.AppendArrayVariant(*t)
	}

	p.SetLastEvaluatedT(arrayT)

	return nil
}
