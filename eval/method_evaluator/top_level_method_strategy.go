package method_evaluator

import (
	"ti/base"
)

type topLevelMethodStrategy struct{}

func (t *topLevelMethodStrategy) evaluate(m *MethodEvaluator) error {
	class, methodT, err := t.getRequiredValues(m)
	if err != nil {
		m.errorResolve()
		return err
	}

	return evaluateNoUnionInstanceMethod(m, class, methodT)
}

func (t *topLevelMethodStrategy) getRequiredValues(m *MethodEvaluator) (
	class string,
	methodT *base.T,
	err error,
) {

	class = m.ctx.GetClass()

	methodT =
		base.GetTopLevelMethodT(m.ctx.GetFrame(), class, m.method)

	if methodT == nil {
		return "", nil, m.makeNotDefinedMethodError("", m.method, "")
	}

	return class, methodT, nil
}
