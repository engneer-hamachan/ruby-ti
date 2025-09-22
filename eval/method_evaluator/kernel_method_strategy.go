package method_evaluator

import (
	"ti/base"
)

func init() {
	dynamicStrategies[[2]string{"Kernel", "yield"}] = &kernelYieldStrategy{}
}

type kernelYieldStrategy struct{}

func (k *kernelYieldStrategy) evaluate(m *MethodEvaluator) error {
	methodT := base.GetMethodT("Builtin", "Kernel", "yield", false)

	if methodT == nil {
		return m.makeNotDefinedMethodError("Kernel", "yield")
	}

	evaluatedArgs, err := getEvaluatedArgs(m, methodT)
	if err != nil {
		return err
	}

	err = checkAndPropagateArgs(m, "Kernel", methodT, evaluatedArgs)
	if err != nil {
		return err
	}

	var blockParameters []base.T

	switch len(evaluatedArgs) {
	case 0:
		blockParameters = append(blockParameters, *base.MakeNil())

	default:
		for _, variant := range evaluatedArgs {
			blockParameters = append(blockParameters, *variant)
		}
	}

	m.parser.SetTmpBlockParameters(blockParameters)
	m.parser.StartParsingExpression()

	m.parser.SetLastEvaluatedT(base.MakeUntyped())

	return nil

}
