package method_evaluator

import (
	"fmt"
	"ti/base"
)

type instanceMethodStrategy struct{}

func (i *instanceMethodStrategy) evaluate(m *MethodEvaluator) error {
	class, methodT, err := i.getRequiredValues(m)

	errorRow := m.parser.ErrorRow
	if err != nil {
		m.errorResolve()
	}

	if err != nil && m.ctx.IsCheckRound() {
		m.parser.ErrorRow = errorRow
		return err
	}

	if methodT == nil {
		return nil
	}

	// protect check
	if methodT.IsProtected {
		callerNode :=
			base.ClassNode{
				Frame: m.ctx.GetFrame(),
				Class: m.ctx.GetClass(),
			}

		methodClassNodes := base.ClassInheritanceMap[callerNode]

		var isContained bool

		if methodT.BeforeEvaluateFrame == m.ctx.GetFrame() && methodT.BeforeEvaluateClass == m.ctx.GetClass() {
			isContained = true
		}

		for _, node := range methodClassNodes {
			if node.Frame == methodT.BeforeEvaluateFrame && node.Class == methodT.BeforeEvaluateClass {
				isContained = true
			}
		}

		if !isContained {
			return fmt.Errorf("%s.%s is protect method", methodT.BeforeEvaluateClass, methodT.GetMethodName())
		}
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

	methodT =
		base.GetMethodT(
			m.evaluatedObjectT.GetFrame(),
			m.evaluatedObjectT.ID,
			m.method,
			false,
		)

	if methodT != nil {
		methodT.SetBeforeEvaluateCode(class + "." + m.method)
		return class, methodT, nil
	}

	if class == "Identifier" {
		return class, methodT, nil
	}

	return "", nil, m.makeNotDefinedMethodError(class, m.method)
}
