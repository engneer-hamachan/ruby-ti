package method_evaluator

import (
	"fmt"
	"ti/base"
)

func init() {
	dynamicStrategies[[2]string{"Kernel", "yield"}] = &kernelYieldStrategy{}
	dynamicStrategies[[2]string{"Kernel", "p"}] = &kernelPrintStrategy{}
}

type kernelYieldStrategy struct{}

func (k *kernelYieldStrategy) evaluate(m *MethodEvaluator) error {
	methodT := base.GetMethodT("Builtin", "Kernel", "yield", false)

	if methodT == nil {
		return m.makeNotDefinedMethodError("Kernel", "yield", "instance")
	}

	base.GlobT = *methodT

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

type kernelPrintStrategy struct{}

func (k *kernelPrintStrategy) evaluate(m *MethodEvaluator) error {
	defineRow := m.parser.ErrorRow

	methodT := base.GetMethodT("Builtin", "Kernel", "p", false)

	if methodT == nil {
		return m.makeNotDefinedMethodError("Kernel", "p", "instance")
	}

	base.GlobT = *methodT

	evaluatedArgs, err := getEvaluatedArgs(m, methodT)
	if err != nil {
		return err
	}

	err = checkAndPropagateArgs(m, "Kernel", methodT, evaluatedArgs)
	if err != nil {
		return err
	}

	var hint string

	hint += "@"
	hint += m.parser.FileName + ":::"
	hint += fmt.Sprintf("%d", defineRow)
	hint += ":::"

	hint += "output: "
	t := m.parser.GetLastEvaluatedT()
	hint += base.TypeToString(&t)

	m.parser.DefineInfos = append(m.parser.DefineInfos, hint)

	switch len(evaluatedArgs) {
	case 0:
		nilT := base.MakeNil()
		m.parser.SetLastEvaluatedT(nilT)

	case 1:
		t := evaluatedArgs[0]
		m.parser.SetLastEvaluatedT(t)

	default:
		arrayT := base.MakeAnyArray()
		for _, variant := range evaluatedArgs {
			arrayT.AppendArrayVariant(*variant)
		}
		m.parser.SetLastEvaluatedT(arrayT)
	}

	return nil
}
