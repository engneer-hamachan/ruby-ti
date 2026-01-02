package method_evaluator

import (
	"ti/base"
)

type unionInstanceStrategy struct{}

func (u *unionInstanceStrategy) evaluate(m *MethodEvaluator) error {
	classNames, methodTs, err := u.getRequiredValues(m)
	if err != nil {
		m.errorResolve()
		return err
	}

	if len(methodTs) == 0 {
		return nil
	}

	evaluatedArgs, err := getEvaluatedArgs(m, methodTs[0])
	if err != nil {
		return err
	}

	return evaluateUnionInstanceMethod(
		m,
		classNames,
		methodTs,
		evaluatedArgs,
	)
}

func (u *unionInstanceStrategy) getRequiredValues(m *MethodEvaluator) (
	classNames []string,
	methodTs []*base.T,
	err error,
) {

	var unionVariants []base.T

	switch m.objectT.ToString() {
	case "union":
		unionVariants = m.objectT.GetVariants()

	default:
		unionVariants = m.evaluatedObjectT.GetVariants()
	}

	for _, t := range unionVariants {
		if t.IsAnyType() {
			continue
		}

		class := t.GetObjectClass()

		methodT :=
			base.GetMethodT(
				m.evaluatedObjectT.GetFrame(),
				class,
				m.method,
				false,
			)

		if methodT != nil {
			methodT.SetBeforeEvaluateCode(class + "." + m.method)
			classNames = append(classNames, class)
			methodTs = append(methodTs, methodT)

			continue
		}

		methodT =
			base.GetInstanceValueT(
				m.evaluatedObjectT.GetFrame(),
				class,
				m.method,
			)

		if methodT != nil {
			methodT.SetBeforeEvaluateCode(class + "." + m.method)
			classNames = append(classNames, class)
			methodTs = append(methodTs, methodT)
			continue
		}

		if m.isAmpersand && class == "NilClass" {
			continue
		}

		err = m.makeNotDefinedMethodError(class, m.method, "instance")

		return []string{}, []*base.T{}, err

	}

	return classNames, methodTs, nil
}
