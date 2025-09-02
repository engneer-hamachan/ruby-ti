package method_evaluator

import (
	"ti/base"
)

type instanceMethodStrategy struct{}

func (i *instanceMethodStrategy) evaluate(m *MethodEvaluator) error {
	class, methodT, err := i.getRequiredValues(m)
	if err != nil && m.ctx.IsCheckRound() {
		return err
	}

	if err != nil {
		m.errorResolve()
	}

	if methodT == nil {
		return nil
	}

	return evaluateNoUnionInstanceMethod(m, class, methodT)
}

func (i *instanceMethodStrategy) getRequiredValues(m *MethodEvaluator) (
	class string,
	methodT *base.T,
	err error,
) {

	class = m.evaluatedObjectT.GetObjectClass()

	methodT =
		base.GetMethodT(m.evaluatedObjectT.GetFrame(), class, m.method, false)

	if methodT != nil {
		methodT.SetBeforeEvaluateCode(class + "." + m.method)
		return class, methodT, nil
	}

	methodT =
		base.GetInstanceValueT(m.evaluatedObjectT.GetFrame(), class, m.method)

	if methodT != nil {
		methodT.SetBeforeEvaluateCode(class + "." + m.method)
		return class, methodT, nil
	}

	if m.objectT.IsAnyType() {
		return "", m.objectT, nil
	}

	if m.isAmpersand && class == "Nil" {
		return class, methodT, nil
	}

	return "", nil, m.makeNotDefinedMethodError(class, m.method)
}
