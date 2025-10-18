package method_evaluator

import (
	"fmt"
	"ti/base"
)

func init() {
	dynamicStrategies[[2]string{"Array", "append"}] = &arrayAppendStrategy{}
	dynamicStrategies[[2]string{"Array", "push"}] = &arrayAppendStrategy{}
	dynamicStrategies[[2]string{"Array", "<<"}] = &arrayAppendStrategy{}
	dynamicStrategies[[2]string{"Array", "concat"}] = &concatArraystrategy{}
	dynamicStrategies[[2]string{"Array", "unshift"}] = &unshiftArraystrategy{}
	dynamicStrategies[[2]string{"Array", "replace"}] = &replaceArraystrategy{}
	dynamicStrategies[[2]string{"Array", "slice"}] = &sliceArrayStrategy{}
	dynamicStrategies[[2]string{"Array", "+"}] = &addArrayStrategy{}
}

type arrayAppendStrategy struct{}

func (a *arrayAppendStrategy) evaluate(m *MethodEvaluator) error {
	methodT := base.GetMethodT("Builtin", "Array", "push", false)

	if methodT == nil {
		return m.makeNotDefinedMethodError("Array", "push")
	}

	evaluatedArgs, err := getEvaluatedArgs(m, methodT)
	if err != nil {
		return err
	}

	err = checkAndPropagateArgs(m, "Array", methodT, evaluatedArgs)
	if err != nil {
		return err
	}

	arrayT := m.evaluatedObjectT

	if len(evaluatedArgs) > 0 {
		arrayT.AppendArrayVariant(*evaluatedArgs[0])
	}

	base.SetValueT(
		m.ctx.GetFrame(),
		m.ctx.GetClass(),
		m.ctx.GetMethod(),
		arrayT.GetBeforeEvaluateCode(),
		arrayT,
	)

	m.parser.SetLastEvaluatedT(arrayT)

	return nil

}

type concatArraystrategy struct{}

func (c *concatArraystrategy) evaluate(m *MethodEvaluator) error {
	methodT := base.GetMethodT("Builtin", "Array", "concat", false)

	if methodT == nil {
		return m.makeNotDefinedMethodError("Array", "concat")
	}

	evaluatedArgs, err := getEvaluatedArgs(m, methodT)
	if err != nil {
		return err
	}

	err = checkAndPropagateArgs(m, "Array", methodT, evaluatedArgs)
	if err != nil {
		return err
	}

	arrayT := m.evaluatedObjectT

	if len(evaluatedArgs) > 0 {
		for _, variant := range evaluatedArgs[0].GetVariants() {
			arrayT.AppendArrayVariant(variant)
		}
	}

	m.parser.SetLastEvaluatedT(arrayT)

	return nil
}

type unshiftArraystrategy struct{}

func (u *unshiftArraystrategy) evaluate(m *MethodEvaluator) error {
	methodT := base.GetMethodT("Builtin", "Array", "unshift", false)

	if methodT == nil {
		return m.makeNotDefinedMethodError("Array", "unshift")
	}

	evaluatedArgs, err := getEvaluatedArgs(m, methodT)
	if err != nil {
		return err
	}

	arrayT := m.evaluatedObjectT

	for _, argT := range evaluatedArgs {
		arrayT.AppendArrayVariant(*argT)
	}

	m.parser.SetLastEvaluatedT(arrayT)

	return nil
}

type replaceArraystrategy struct{}

func (r *replaceArraystrategy) evaluate(m *MethodEvaluator) error {
	methodT := base.GetMethodT("Builtin", "Array", "replace", false)

	if methodT == nil {
		return m.makeNotDefinedMethodError("Array", "replace")
	}

	evaluatedArgs, err := getEvaluatedArgs(m, methodT)
	if err != nil {
		return err
	}

	err = checkAndPropagateArgs(m, "Array", methodT, evaluatedArgs)
	if err != nil {
		return err
	}

	newArrayT := evaluatedArgs[0]
	newArrayT.SetBeforeEvaluateCode(m.evaluatedObjectT.GetBeforeEvaluateCode())

	m.parser.SetLastEvaluatedT(newArrayT)

	base.SetValueT(
		m.ctx.GetFrame(),
		m.ctx.GetClass(),
		m.ctx.GetMethod(),
		m.objectT.ToString(),
		newArrayT,
	)

	return nil
}

type sliceArrayStrategy struct{}

func (s *sliceArrayStrategy) evaluate(m *MethodEvaluator) error {
	methodT := base.GetMethodT("Builtin", "Array", "slice", false)

	if methodT == nil {
		return m.makeNotDefinedMethodError("Array", "slice")
	}

	arrayT := m.evaluatedObjectT

	m.parser.SetLastEvaluatedT(arrayT.UnifyVariants())

	if !arrayT.IsArrayType() {
		return fmt.Errorf("%s is not Array", m.objectT.ToString())
	}

	nextT, isParentheses, err := m.parser.ReadWithCheck("(")
	if err != nil {
		return err
	}

	if isParentheses {
		nextT, err = m.parser.Read()
		if err != nil {
			return err
		}

		defer m.parser.Skip()
	}

	err = m.outerEval.Eval(m.parser, m.ctx, nextT)
	if err != nil {
		return err
	}

	idxT := m.parser.GetLastEvaluatedT()
	if idxT.GetType() != base.INT {
		return fmt.Errorf("%s is not Integer", nextT.ToString())
	}

	unionT := arrayT.UnifyVariants()
	m.parser.SetLastEvaluatedT(unionT)

	base.SetValueT(
		m.ctx.GetFrame(),
		m.ctx.GetClass(),
		m.ctx.GetMethod(),
		m.objectT.ToString(),
		unionT,
	)

	return nil
}

type addArrayStrategy struct{}

func (a *addArrayStrategy) evaluate(m *MethodEvaluator) error {
	methodT := base.GetMethodT("Builtin", "Array", "+", false)

	if methodT == nil {
		return m.makeNotDefinedMethodError("Array", "+")
	}

	evaluatedArgs, err := getEvaluatedArgs(m, methodT)
	if err != nil {
		return err
	}

	err = checkAndPropagateArgs(m, "Array", methodT, evaluatedArgs)
	if err != nil {
		return err
	}

	arrayT := m.evaluatedObjectT
	argT := evaluatedArgs[0]

	for _, variant := range argT.GetVariants() {
		arrayT.AppendArrayVariant(variant)
	}

	m.parser.SetLastEvaluatedT(arrayT)

	return nil
}
