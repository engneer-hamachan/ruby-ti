package method_evaluator

import (
	"fmt"
	"slices"
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

	if m.parser.IsStrict && m.ctx.IsCheckRound() && class == "Identifier" {
		return m.makeNotDefinedMethodError(class, m.method, "instance")
	}

	if methodT == nil {
		return nil
	}

	if m.parser.LspTargetRow == m.parser.ErrorRow {
		base.GlobT = *methodT
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

		if methodT.DefinedFrame == m.ctx.GetFrame() && methodT.DefinedClass == m.ctx.GetClass() {
			isContained = true
		}

		if !isContained {
			for _, node := range methodClassNodes {
				if node.Frame == methodT.DefinedFrame && node.Class == methodT.DefinedClass {
					isContained = true
				}
			}
		}

		if !isContained {
			return fmt.Errorf("%s.%s is protect method", methodT.DefinedClass, methodT.GetMethodName())
		}
	}

	return evaluateNoUnionInstanceMethod(m, class, methodT)
}

func (i *instanceMethodStrategy) isTransformIdentifier(method string) bool {
	transformTargetIdentifiers := []string{
		"%",
		"&",
		"*",
		"**",
		"+",
		"+@",
		"-",
		"-@",
		"/",
		"<",
		"<<",
		"<=",
		"<=>",
		"==",
		"===",
		">",
		">=",
		">>",
		"^",
		"|",
	}

	return slices.Contains(transformTargetIdentifiers, method)
}

func (i *instanceMethodStrategy) isResolveMethodCall(
	m *MethodEvaluator,
	class string,
) bool {

	if m.ctx.IsCollectRound() {
		return false
	}

	targetClasses := []string{"Untyped", "Identifier"}

	if !slices.Contains(targetClasses, class) {
		return false
	}

	return i.isTransformIdentifier(m.method)
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

	if i.isResolveMethodCall(m, class) {
		methodT =
			base.MakeMethod(
				class,
				m.method,
				*base.MakeUntyped(),
				[]string{base.GenId()},
			)

		return class, methodT, nil
	}

	if m.objectT.IsAnyType() {
		return "", m.objectT, nil
	}

	if m.isAmpersand && class == "NilClass" {
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

	if class == "Untyped" || class == "Identifier" {
		return class, base.MakeUntyped(), nil
	}

	return "", nil, m.makeNotDefinedMethodError(class, m.method, "instance")
}
