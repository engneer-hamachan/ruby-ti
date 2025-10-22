package method_evaluator

import (
	"ti/base"
)

func handleRefference(m *MethodEvaluator) error {
	nextT, err := m.parser.Read()
	if err != nil {
		return err
	}

	if nextT.IsTargetIdentifier("[") {
		m.ctx.IsBind = false
		m.outerEval.Eval(m.parser, m.ctx, nextT)

		return nil
	}

	m.parser.Unget()

	return nil
}

func removeBlockTypeArgs(argTs []*base.T) []*base.T {
	var variants []*base.T

	for _, argT := range argTs {
		if argT.GetType() != base.BLOCK {
			variants = append(variants, argT)
		}
	}

	return variants
}

func conditioningMethodReturn(
	m *MethodEvaluator,
	class string,
	methodT *base.T,
	evaluatedArgs []*base.T,
) *base.T {

	defineArgTs := getEvaluatedDefineArgs(m, class, methodT)

	for _, defineArgT := range defineArgTs {
		if defineArgT.HasDefault() {
			variants := methodT.GetVariants()
			return &variants[len(removeBlockTypeArgs(evaluatedArgs))]
		}

		if defineArgT.IsUnionType() {
			for idx, variant := range defineArgT.GetVariants() {
				for _, argT := range evaluatedArgs {
					if variant.GetType() == argT.GetType() || (variant.IsAnyType()) {
						variants := methodT.GetVariants()
						return &variants[idx]
					}
				}
			}

			continue
		}

		for idx, argT := range evaluatedArgs {
			if defineArgT.GetType() == argT.GetType() || (defineArgT.IsAnyType()) {
				variants := methodT.GetVariants()
				return &variants[idx]
			}
		}
	}

	return methodT
}

func setLastEvaluatedTAndReturn(
	m *MethodEvaluator,
	methodT *base.T,
	args []*base.T,
) *base.T {

	switch methodT.GetType() {
	case base.BLOCK:
		m.parser.SetLastEvaluatedT(methodT.GetVal())
		return methodT.GetVal().(*base.T)

	case base.UNION:
		var newVariants []base.T
		for _, variant := range methodT.GetVariants() {
			if variant.GetType() == base.UNIFY {
				unifiedT := m.evaluatedObjectT.UnifyVariants()
				newVariants = append(newVariants, *unifiedT)
				continue
			}

			newVariants = append(newVariants, variant)
		}

		unionT := base.MakeUnifiedT(newVariants)
		m.parser.SetLastEvaluatedT(unionT)

		return unionT

	case base.SELF:
		m.parser.SetLastEvaluatedT(m.evaluatedObjectT)
		return m.evaluatedObjectT

	case base.SELF_CONVERT_ARRAY:
		arrayT := base.MakeAnyArray()

		for _, variant := range m.evaluatedObjectT.GetVariants() {
			arrayT.AppendArrayVariant(variant)
		}

		m.parser.SetLastEvaluatedT(arrayT)

		return arrayT

	case base.SELF_ARGUMENT:
		switch len(args) {
		case 0:
			nilT := base.MakeNil()
			m.parser.SetLastEvaluatedT(nilT)
			return nilT

		case 1:
			t := args[0]
			m.parser.SetLastEvaluatedT(t)
			return t

		default:
			arrayT := base.MakeAnyArray()
			for _, variant := range args {
				arrayT.AppendArrayVariant(*variant)
			}
			m.parser.SetLastEvaluatedT(arrayT)
			return arrayT
		}

	case base.UNIFY:
		unifiedT := m.evaluatedObjectT.UnifyVariants()
		m.parser.SetLastEvaluatedT(unifiedT)

		return unifiedT

	case base.OPTIONAL_UNIFY:
		m.evaluatedObjectT.AppendVariant(*base.MakeNil())

		unifiedT := base.MakeUnifiedT(m.evaluatedObjectT.GetVariants())
		m.parser.SetLastEvaluatedT(unifiedT)

		return unifiedT

	case base.BLOCK_RESULT_ARRAY:
		blockT := m.parser.GetLastEvaluatedT()
		blockResultT := blockT.GetVal().(*base.T)

		arrayT := base.MakeArray()
		arrayT.AppendArrayVariant(*blockResultT)

		m.parser.SetLastEvaluatedT(arrayT)

		return arrayT

	case base.KEYVALUE_ARRAY:
		hashT := m.evaluatedObjectT
		arrayT := base.MakeAnyArray()

		for _, variants := range hashT.GetVariants() {
			arrayT.AppendArrayVariant(*variants.GetKeyValue())
		}

		m.parser.SetLastEvaluatedT(arrayT)
		return arrayT

	default:
		if methodT.IsNameSpaceIdentifier() {
			frame, parentClass, class :=
				base.SeparateNameSpaces(methodT.ToString())

			t := base.MakeObject(class)
			t.SetFrame(base.CalculateFrame(frame, parentClass))
			methodT = t
		}

		m.parser.SetLastEvaluatedT(methodT)
		return methodT
	}
}

func evaluateNoUnionInstanceMethod(
	m *MethodEvaluator,
	class string,
	methodT *base.T,
) error {

	evaluatedArgs, err := getEvaluatedArgs(m, methodT)
	if err != nil {
		return err
	}

	err = checkAndPropagateArgs(m, class, methodT, evaluatedArgs)
	if err != nil {
		return err
	}

	if methodT.IsConditionalReturn {
		methodT = conditioningMethodReturn(m, class, methodT, evaluatedArgs)
	}

	returnT := setLastEvaluatedTAndReturn(m, methodT, evaluatedArgs)

	// a&.b
	if m.isAmpersand {
		switch returnT.IsUnionType() {
		case true:
			variants := returnT.GetVariants()
			variants = append(variants, *base.MakeNil())
			returnT = base.MakeUnion(variants)
		default:
			variants := []base.T{*returnT, *base.MakeNil()}
			returnT = base.MakeUnion(variants)
		}
	}

	if methodT.IsDestructive {
		base.SetValueT(
			m.ctx.GetFrame(),
			m.ctx.GetClass(),
			m.ctx.GetMethod(),
			m.objectT.ToString(),
			returnT,
			m.ctx.IsDefineStatic,
		)
	}

	return handleRefference(m)
}

func evaluateUnionInstanceMethod(
	m *MethodEvaluator,
	classNames []string,
	methodTs []*base.T,
	evaluatedArgs []*base.T,
) error {

	returnT, err :=
		checkAndPropagateArgsForUnionWithReturnT(
			m,
			classNames,
			methodTs,
			evaluatedArgs,
		)

	if err != nil {
		return err
	}

	// a&.b
	if m.isAmpersand {
		switch returnT.IsUnionType() {
		case true:
			variants := returnT.GetVariants()
			variants = append(variants, *base.MakeNil())
			returnT = base.MakeUnion(variants)
		default:
			variants := []base.T{*returnT, *base.MakeNil()}
			returnT = base.MakeUnion(variants)
		}
	}

	m.parser.SetLastEvaluatedT(returnT)

	return handleRefference(m)
}
