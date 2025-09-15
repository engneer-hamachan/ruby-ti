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

	frame := m.ctx.GetFrame()

	if m.ctx.GetClass() != "" {
		switch frame {
		case "":
			frame = m.ctx.GetClass()
		default:
			frame = frame + "::" + m.ctx.GetClass()
		}
	}

	methodT =
		base.GetClassMethodT(frame, class, m.method, false)

	if methodT == nil {
		return "", nil, m.makeNotDefinedMethodError(class, m.method)
	}

	methodT.SetBeforeEvaluateCode(class + "." + m.method)

	return class, methodT, nil
}
