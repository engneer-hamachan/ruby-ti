package eval

import (
	"fmt"
	"ti/base"
	"ti/context"
	"ti/parser"
)

type IfUnless struct {
	conditionType string
	originalTs    map[string][]base.T
	narrowTs      map[string][]base.T
}

func NewIfUnless(conditionType string) DynamicEvaluator {
	return &IfUnless{conditionType: conditionType}
}

func init() {
	DynamicEvaluators["if"] = NewIfUnless("if")
	DynamicEvaluators["unless"] = NewIfUnless("unless")
}

func (i *IfUnless) isSpecialCtxMethod(t *base.T) bool {
	specialMethods := []string{"is_a?", "=="}

	return t.IsTargetIdentifiers(specialMethods)
}

func (i *IfUnless) convertClassNameToTobject(class string) *base.T {
	switch class {
	case "String":
		return base.MakeAnyString()
	case "Integer":
		return base.MakeAnyInt()
	case "Float":
		return base.MakeAnyFloat()
	case "NilClass":
		return base.MakeNil()
	case "Hash":
		return base.MakeAnyHash()
	case "Array":
		return base.MakeAnyArray()
	case "Bool":
		return base.MakeBool()
	}

	return base.MakeObject(class)
}

func (i *IfUnless) setConditionalCtx(
	class string,
	object string,
	ctx context.Context,
	currentT base.T,
	isExclamation bool,
) error {

	classT := i.convertClassNameToTobject(class)

	var isNarrow bool

	switch i.conditionType {
	case "if":
		isNarrow = !isExclamation

	case "unless":
		isNarrow = isExclamation

	default:
		return fmt.Errorf("syntax error")
	}

	i.originalTs[object] = append(i.originalTs[object], currentT)

	switch isNarrow {
	case true:
		base.SetValueT(
			ctx.GetFrame(),
			ctx.GetClass(),
			ctx.GetMethod(),
			object,
			classT,
			ctx.IsDefineStatic,
		)

		i.narrowTs[object] = append(i.narrowTs[object], *classT)

	default:
		if currentT.IsUnionType() {
			var newVariants []base.T
			for _, currentVariant := range currentT.GetVariants() {
				if currentVariant.GetObjectClass() != classT.GetObjectClass() {
					newVariants = append(newVariants, currentVariant)
				}
			}

			base.SetValueT(
				ctx.GetFrame(),
				ctx.GetClass(),
				ctx.GetMethod(),
				object,
				base.MakeUnifiedT(newVariants),
				ctx.IsDefineStatic,
			)

			i.narrowTs[object] = newVariants

			break
		}

		if currentT.IsMatchType(classT) {
			base.SetValueT(
				ctx.GetFrame(),
				ctx.GetClass(),
				ctx.GetMethod(),
				object,
				base.MakeNil(),
				ctx.IsDefineStatic,
			)

			i.narrowTs[object] = append(i.narrowTs[object], *base.MakeNil())
		}
	}

	return nil
}

func (i *IfUnless) beforeEval(
	e Evaluator,
	p parser.Parser,
	ctx context.Context,
) error {
	nextT, err := p.Read()
	if err != nil {
		p.Fatal(ctx, err)
	}

	err = e.Eval(&p, ctx, nextT)
	if err != nil {
		p.Fatal(ctx, err)
	}

	return nil
}

func (i *IfUnless) getBackupContext(
	e *Evaluator,
	p parser.Parser,
	ctx context.Context,
) (func(), error) {

	err := i.beforeEval(*e, p, ctx)
	if err != nil {
		return func() {}, err
	}

	var isExclamation bool

	nextT, err := p.Read()
	if err != nil {
		return func() {}, err
	}

	if nextT.IsExclamationIdentifier() {
		isExclamation = true

		nextT, err = p.Read()
		if err != nil {
			return func() {}, err
		}
	}

	if err != nil || !nextT.IsIdentifierType() {
		return func() {}, err
	}

	var class string
	object := nextT.ToString()

	nextT, err = p.Read()
	if err != nil {
		return func() {}, err
	}

	if nextT.IsDotIdentifier() {
		nextT, err = p.Read()
		if err != nil {
			return func() {}, err
		}
	}

	if !i.isSpecialCtxMethod(nextT) {
		return func() {}, err
	}

	nextT, err = p.Read()
	if err != nil {
		return func() {}, err
	}

	switch nextT.IsOpenParentheses() {
	case true:
		nextT, err = p.Read()
		if err != nil {
			return func() {}, err
		}

		if !nextT.IsClassType() {
			return func() {}, fmt.Errorf(
				"type missmatch error: given %s expected Class for is_a?",
				base.TypeToString(nextT),
			)
		}

		class = nextT.ToString()

		nextT, err = p.Read()
		if err != nil || !nextT.IsIdentifierType() {
			return func() {}, err
		}

	default:
		class = nextT.GetObjectClass()

		nextT, err = p.Read()
		if err != nil || !nextT.IsIdentifierType() {
			return func() {}, err
		}
	}

	t :=
		base.GetDynamicValueT(
			ctx.GetFrame(),
			ctx.GetClass(),
			ctx.GetMethod(),
			object,
		)

	if t == nil {
		return func() {}, nil
	}

	err = i.setConditionalCtx(class, object, ctx, *t, isExclamation)
	if err != nil {
		return func() {}, err
	}

	return func() {
		base.SetValueT(
			ctx.GetFrame(),
			ctx.GetClass(),
			ctx.GetMethod(),
			object,
			t,
			ctx.IsDefineStatic,
		)
	}, nil
}

func (i *IfUnless) getEndIdentifier(p *parser.Parser) string {
	switch p.IsParsingExpression() {
	case true:
		return "\n"

	default:
		return "end"
	}
}

func (i *IfUnless) Evaluation(
	e *Evaluator,
	p *parser.Parser,
	ctx context.Context,
	t *base.T,
) (err error) {

	// clear
	i.originalTs = make(map[string][]base.T)
	i.narrowTs = make(map[string][]base.T)

	lastEvaluatedT := p.GetLastEvaluatedT()

	zaorik, err := i.getBackupContext(e, *p, ctx)
	if err != nil {
		p.Fatal(ctx, err)
	}

	defer zaorik()

	endIdentifier := i.getEndIdentifier(p)

	nextT, err := p.Read()
	if err != nil {
		p.Fatal(ctx, err)
	}

	err = e.Eval(p, ctx, nextT)
	if err != nil {
		p.Fatal(ctx, err)
	}

	resultTs := []base.T{*base.MakeNil()}

	for {
		nextT, err := p.Read()
		if err != nil {
			p.Fatal(ctx, err)
		}

		if nextT == nil {
			return nil
		}

		if nextT.IsTargetIdentifier(endIdentifier) {
			if !p.IsParsingExpression() {
				resultTs = append(resultTs, p.GetLastEvaluatedT())
			}

			break
		}

		if nextT.IsTargetIdentifier("elsif") && i.conditionType == "if" {
			_, err := i.getBackupContext(e, *p, ctx)
			if err != nil {
				p.Fatal(ctx, err)
			}

			resultTs = append(resultTs, p.GetLastEvaluatedT())
			continue
		}

		if nextT.IsTargetIdentifier("else") {
			p.SkipNewline()

			nextT, err := p.Read()
			if err != nil {
				p.Fatal(ctx, err)
			}

			p.Unget()

			if nextT.ToString() != endIdentifier {
				resultTs = resultTs[1:]
			}

			resultTs = append(resultTs, p.GetLastEvaluatedT())

			// narrowing proccess
			for originalKey, originalVariants := range i.originalTs {
				narrowVariants, ok := i.narrowTs[originalKey]
				if !ok {
					continue
				}

				variants := []base.T{}

				for _, originalVariant := range originalVariants {
					switch originalVariant.GetType() {
					case base.UNION:
						for _, variant := range originalVariant.GetVariants() {
							isContain := false

							for _, narrowVariant := range narrowVariants {
								if variant.GetObjectClass() == narrowVariant.GetObjectClass() {
									isContain = true
								}
							}

							if isContain {
								continue
							}

							variants = append(variants, variant)
						}

					default:
						for _, narrowVariant := range narrowVariants {
							if originalVariant.IsEqualObject(&narrowVariant) &&
								originalVariant.GetObjectClass() != narrowVariant.GetObjectClass() {

								continue
							}

							variants = append(variants, originalVariant)
						}
					}
				}

				base.SetValueT(
					ctx.GetFrame(),
					ctx.GetClass(),
					ctx.GetMethod(),
					originalKey,
					base.MakeUnifiedT(variants),
					ctx.IsDefineStatic,
				)
			}

			continue
		}

		err = e.Eval(p, ctx, nextT)
		if err != nil {
			p.Fatal(ctx, err)
		}
	}

	if endIdentifier == "\n" {
		p.Unget()
	}

	if p.IsParsingExpression() {
		resultTs = append(resultTs, lastEvaluatedT)
	}

	resultT := base.MakeUnifiedT(resultTs)
	e.setLastEvaluatedT(p, ctx, resultT)

	return nil
}
