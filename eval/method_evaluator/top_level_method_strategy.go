package method_evaluator

import (
	"ti/base"
	"ti/parser"
)

type topLevelMethodStrategy struct{}

func (t *topLevelMethodStrategy) evaluate(m *MethodEvaluator) error {
	class, methodT, err := t.getRequiredValues(m)
	if err != nil {
		m.errorResolve()
		return err
	}

	if m.parser.LspTargetRow == m.parser.ErrorRow {
		parser.GlobT = *methodT
	}

	return evaluateNoUnionInstanceMethod(m, class, methodT)
}

func (t *topLevelMethodStrategy) getRequiredValues(m *MethodEvaluator) (
	class string,
	methodT *base.T,
	err error,
) {

	class = m.ctx.GetClass()

	switch m.ctx.IsDefineStatic {
	case true:
		// TODO: refact start
		methodT =
			base.GetTopLevelClassMethodT(m.ctx.GetFrame(), class, m.method)

		if methodT == nil {
			methodT =
				base.GetTopLevelMethodT(m.ctx.GetFrame(), class, m.method)
		}
		// TODO: refact end
	default:
		methodT =
			base.GetTopLevelMethodT(m.ctx.GetFrame(), class, m.method)
	}

	if methodT == nil {
		return "", nil, m.makeNotDefinedMethodError("", m.method, "")
	}

	return class, methodT, nil
}
