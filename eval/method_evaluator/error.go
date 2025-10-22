package method_evaluator

import (
	"fmt"
	"ti/base"
)

func makeTypeError(
	class, method string,
	currentTypeStr string,
	wantTypeStr string,
) (err error) {

	var code string

	switch class {
	case "":
		code = method
	default:
		code = class + "." + method
	}

	return fmt.Errorf(
		"type mismatch: expected %s, but got %s for %s",
		wantTypeStr,
		currentTypeStr,
		code,
	)
}

func (m *MethodEvaluator) makeNotDefinedMethodError(
	class, method string,
) error {

	if class == "" {
		return fmt.Errorf("method '%s' is not defined", method)
	}

	return fmt.Errorf("method '%s' is not defined for %s", method, class)
}

func (m *MethodEvaluator) makeNotDefinedClassError(class string) error {
	return fmt.Errorf("class '%s' is not defined", class)
}

func makeDefineArgumentInfo(
	m *MethodEvaluator,
	class string,
	methodT *base.T,
) string {

	argumentTypes := "("

	for _, definedArg := range methodT.GetDefineArgs() {
		if argumentTypes != "(" {
			argumentTypes += ", "
		}

		if isKeySuffix(definedArg) {
			argumentTypes += definedArg + " "
			definedArg = removeSuffix(definedArg)
		}

		if isAsteriskPrefix(definedArg) {
			argumentTypes += "*"
		}

		definedArgT :=
			base.GetValueT(
				methodT.GetFrame(),
				class,
				m.method,
				definedArg,
				methodT.IsStatic,
			)

		if definedArgT.HasDefault() {
			argumentTypes += "?"
		}

		if isAsteriskPrefix(definedArg) {
			definedArgT = definedArgT.UnifyVariants()
		}

		switch definedArgT.GetType() {
		case base.UNION:
			argumentTypes += base.UnionTypeToString(definedArgT.GetVariants())

		case base.UNKNOWN:
			argumentTypes += "?"

		default:
			argumentTypes += base.TypeToString(definedArgT)
		}
	}

	argumentTypes += ")"

	return argumentTypes
}

func (m *MethodEvaluator) errorResolve() error {
	for {
		nextT, err := m.parser.Read()
		if err != nil {
			return err
		}

		if nextT.IsTargetIdentifier("{") {
			m.parser.Unget()
			_, err = expectBlockArgProcess(m, base.MakeIdentifier("tmp"), []*base.T{})
			return err
		}

		if nextT == nil || nextT.IsNewLineIdentifier() {
			m.parser.Unget()
			break
		}

		err = m.outerEval.Eval(m.parser, m.ctx, nextT)
		if err != nil {
			return err
		}
	}

	m.parser.ConsumeLastReturnT()
	m.parser.SetLastEvaluatedT(base.MakeUnknown())

	return nil
}
