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

	if m.parser.LspTargetRow == m.parser.ErrorRow {
		base.GlobT = *methodT
	}

	return evaluateNoUnionInstanceMethod(m, class, methodT)
}

func (t *topLevelMethodStrategy) isMismatchVisibility(
	m *MethodEvaluator,
	methodT *base.T,
) bool {

	return m.ctx.IsDefineStatic != methodT.IsStatic &&
		m.ctx.GetMethod() != "" &&
		m.ctx.GetMethod() != "new" &&
		methodT.GetFrame() != "Builtin"
}

func (t *topLevelMethodStrategy) getRequiredValues(m *MethodEvaluator) (
	class string,
	methodT *base.T,
	err error,
) {

	class = m.ctx.GetClass()

	switch m.ctx.IsDefineStatic {
	case true:
		methodT = base.GetTopLevelClassMethodT(m.ctx.GetFrame(), class, m.method)

		if methodT == nil {
			methodT = base.GetTopLevelMethodT(m.ctx.GetFrame(), class, m.method)
		}

	default:
		methodT = base.GetTopLevelMethodT(m.ctx.GetFrame(), class, m.method)
	}

	if t.isMismatchVisibility(m, methodT) {
		return "", nil, m.makeNotDefinedMethodError("", m.method, "")
	}

	if methodT == nil {
		return "", nil, m.makeNotDefinedMethodError("", m.method, "")
	}

	return class, methodT, nil
}
