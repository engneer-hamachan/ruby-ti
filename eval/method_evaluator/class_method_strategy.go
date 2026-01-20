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

	if m.parser.LspTargetRow == m.parser.ErrorRow {
		base.GlobT = *methodT
	}

	return evaluateNoUnionInstanceMethod(m, class, methodT)
}

func (c *classMethodStrategy) getRequiredValues(m *MethodEvaluator) (
	class string,
	methodT *base.T,
	err error,
) {

	switch m.evaluatedObjectT.IsClassType() {
	case true:
		class = m.evaluatedObjectT.ToString()

	default:
		class = m.objectT.ToString()
	}

	// GetClassMethod with nested scope lookup.
	calculatedFrame := m.ctx.GetFrame()

	if m.ctx.GetClass() != "" {
		switch calculatedFrame {
		case "":
			calculatedFrame = m.ctx.GetClass()
		default:
			calculatedFrame = calculatedFrame + "::" + m.ctx.GetClass()
		}
	}

	frames := []string{calculatedFrame, m.ctx.GetFrame(), ""}
	if !base.IsClassDefined(frames, class) {
		return "", nil, m.makeNotDefinedClassError(class)
	}

	methodT = base.GetClassMethodT(calculatedFrame, class, m.method, false)

	if methodT == nil {
		methodT =
			base.GetClassMethodT(m.ctx.GetFrame(), class, m.method, false)
	}

	if methodT == nil {
		methodT = base.GetClassMethodT("", class, m.method, false)
	}

	if methodT == nil {
		return "", nil, m.makeNotDefinedMethodError(class, m.method, "class")
	}
	// end

	methodT.SetBeforeEvaluateCode(class + "." + m.method)

	return class, methodT, nil
}
