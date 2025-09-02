package method_evaluator

import (
	"ti/base"
)

type classMethodStrategy struct{}

func (c *classMethodStrategy) evaluate(m *MethodEvaluator) error {
	class, methodT, err := c.getRequiredValues(m)
	if err != nil {
		m.errorResolve()
		return err
	}

	return evaluateNoUnionInstanceMethod(m, class, methodT)
}

func (c *classMethodStrategy) getRequiredValues(m *MethodEvaluator) (
	class string,
	methodT *base.T,
	err error,
) {

	class = m.objectT.ToString()

	methodT =
		base.GetClassMethodT(m.ctx.GetFrame(), class, m.method, false)

	if methodT == nil {
		return "", nil, m.makeNotDefinedMethodError(class, m.method)
	}

	methodT.SetBeforeEvaluateCode(class + "." + m.method)

	return class, methodT, nil
}
