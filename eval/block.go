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
				ctx.IsDefineStatic,
			)
		}
	}
}

func calculateBlockParameterType(paramT base.T, evaluatedObjectT *base.T) base.T {
	switch paramT.GetType() {
	case base.UNIFY:
		return *evaluatedObjectT.UnifyVariants()

	case base.ARRAY:
		// Recursively process array inner types
		var newVariants []base.T
		for _, variant := range paramT.GetVariants() {
			processedT := calculateBlockParameterType(variant, evaluatedObjectT)
			newVariants = append(newVariants, processedT)
		}

		arrayT := base.MakeAnyArray()
		for _, variant := range newVariants {
			arrayT.AppendArrayVariant(variant)
		}
		return *arrayT

	case base.UNION:
		// Recursively process union variants
		var newVariants []base.T
		for _, variant := range paramT.GetVariants() {
			processedT := calculateBlockParameterType(variant, evaluatedObjectT)
			newVariants = append(newVariants, processedT)
		}
		return *base.MakeUnifiedT(newVariants)

	default:
		return paramT
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

	if methodT == nil {
		methodT = base.GetClassMethodT(frame, class, method, false)
	}

	if methodT != nil {
		var blockParamaters []base.T

		for _, t := range methodT.GetBlockParameters() {
			// Handle ARRAY and UNION with inner Unify
			if t.GetType() == base.ARRAY || t.GetType() == base.UNION {
				processedT := calculateBlockParameterType(t, &lastEvaluatedT)
				blockParamaters = append(blockParamaters, processedT)
				continue
			}

			if t.GetType() == base.UNIFY {
				blockParamaters =
					append(blockParamaters, *lastEvaluatedT.UnifyVariants())

				continue
			}

			if t.GetType() == base.FLATTEN {
				tmpParameters := [20]*base.T{}

				if len(lastEvaluatedT.UnifyVariants().GetVariants()) == 0 {
					blockParamaters =
						append(blockParamaters, *lastEvaluatedT.UnifyVariants())

					continue
				}

				for idx, variant := range lastEvaluatedT.UnifyVariants().GetVariants() {
					switch variant.GetType() {
					case base.ARRAY:
						arrayVariants := variant.GetVariants()

						for idx, arrayVariant := range arrayVariants {
							if tmpParameters[idx] == nil {
								arrayT := base.MakeAnyArray()
								tmpParameters[idx] = arrayT
							}

							tmpParameters[idx].AppendArrayVariant(arrayVariant)
						}

					case base.KEYVALUE:
						hashT := base.MakeAnyHash()
						hashT.AppendHashVariant(variant)
						tmpParameters[idx] = hashT
						continue

					case base.OBJECT:
						if tmpParameters[idx] == nil {
							arrayT := base.MakeAnyArray()
							tmpParameters[idx] = arrayT
						}

						tmpParameters[idx].AppendArrayVariant(variant)

					default:
						if tmpParameters[0] == nil {
							unionT := base.MakeUnion([]base.T{variant})
							tmpParameters[0] = unionT
							continue
						}

						tmpParameters[0].AppendVariant(variant)
					}
				}

				// max
				var maxLength int
				for _, variant := range tmpParameters {
					if variant != nil && len(variant.GetVariants()) > maxLength {
						maxLength = len(variant.GetVariants())
					}
				}

				var newParameters []base.T

				for _, variant := range tmpParameters {
					if variant == nil {
						continue
					}

					if variant.IsArrayType() && len(variant.GetVariants()) < maxLength {
						variant.AppendArrayVariant(*base.MakeNil())
					}

					if variant.IsHashType() {
						newParameters = append(newParameters, *variant)
						continue
					}

					switch len(variant.GetVariants()) {
					case 0:
						newParameters = append(newParameters, *variant)
					default:
						newParameters = append(newParameters, *variant.UnifyVariants())
					}
				}

				blockParamaters = newParameters

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

			if t.IsNameSpaceIdentifier() {
				frame, parentClass, class := base.SeparateNameSpaces(t.ToString())
				t = *base.MakeObject(class)
				t.SetFrame(base.CalculateFrame(frame, parentClass))
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
				ctx.IsDefineStatic,
			)

			continue
		}

		base.SetValueT(
			ctx.GetFrame(),
			ctx.GetClass(),
			ctx.GetMethod(),
			variable.ToString(),
			&blockParameters[idx],
			ctx.IsDefineStatic,
		)
	}
}

func (d *Do) collectBlockVariables(p *parser.Parser) ([]*base.T, error) {
	_, ok, err := p.ReadWithCheck("|")
	if err != nil {
		return []*base.T{}, err
	}

	if !ok {
		p.Unget()

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

func (d *Do) prepareBlockScope(p *parser.Parser, ctx context.Context) (func(), error) {
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
				ctx.IsDefineStatic,
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

	zaorik, err := d.prepareBlockScope(p, ctx)
	if err != nil {
		return err
	}

	defer zaorik()

	for {
		nextT, err := p.Read()
		if err != nil {
			return err
		}

		if nextT == nil {
			break
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
