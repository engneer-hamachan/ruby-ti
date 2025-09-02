package method_evaluator

import (
	"ti/base"
)

func init() {
	dynamicStrategies[[2]string{"Hash", "merge"}] = &hashMergeStrategy{}

	dynamicStrategies[[2]string{"Hash", "merge!"}] =
		&hashDestructionMergeStrategy{}

	dynamicStrategies[[2]string{"Hash", "shift"}] = &hashShiftStrategy{}
}

type hashMergeStrategy struct{}

func (h *hashMergeStrategy) evaluate(m *MethodEvaluator) error {
	methodT := base.GetMethodT("Builtin", "Hash", "merge", false)

	if methodT == nil {
		return m.makeNotDefinedMethodError("Hash", "merge")
	}

	evaluatedArgs, err := getEvaluatedArgs(m, methodT)
	if err != nil {
		return err
	}

	err = checkAndPropagateArgs(m, "Hash", methodT, evaluatedArgs)
	if err != nil {
		return err
	}

	hashT := m.evaluatedObjectT.DeepCopy()

	hashT.MergeHash(evaluatedArgs[0])

	m.parser.SetLastEvaluatedT(hashT)

	return nil

}

type hashDestructionMergeStrategy struct{}

func (h *hashDestructionMergeStrategy) evaluate(m *MethodEvaluator) error {
	methodT := base.GetMethodT("Builtin", "Hash", "merge!", false)

	if methodT == nil {
		return m.makeNotDefinedMethodError("Hash", "merge!")
	}

	evaluatedArgs, err := getEvaluatedArgs(m, methodT)
	if err != nil {
		return err
	}

	err = checkAndPropagateArgs(m, "Hash", methodT, evaluatedArgs)
	if err != nil {
		return err
	}

	hashT := m.evaluatedObjectT

	hashT.MergeHash(evaluatedArgs[0])

	m.parser.SetLastEvaluatedT(hashT)

	return nil

}

type hashShiftStrategy struct{}

func (h *hashShiftStrategy) evaluate(m *MethodEvaluator) error {
	methodT := base.GetMethodT("Builtin", "Hash", "shift", false)

	if methodT == nil {
		return m.makeNotDefinedMethodError("Hash", "shift")
	}

	evaluatedArgs, err := getEvaluatedArgs(m, methodT)
	if err != nil {
		return err
	}

	err = checkAndPropagateArgs(m, "Hash", methodT, evaluatedArgs)
	if err != nil {
		return err
	}

	hashT := m.evaluatedObjectT
	arrayT := base.MakeAnyArray()

	variantT := hashT.PopVariants()

	if variantT != nil {
		keyT :=
			base.MakeUnion([]base.T{*base.MakeAnyString(), *base.MakeAnySymbol()})

		arrayT.AppendArrayVariant(*keyT)
		arrayT.AppendArrayVariant(*variantT.GetKeyValue())
	}

	base.SetValueT(
		m.ctx.GetFrame(),
		m.ctx.GetClass(),
		m.ctx.GetMethod(),
		m.objectT.ToString(),
		hashT,
	)

	m.parser.SetLastEvaluatedT(arrayT)

	return nil

}
