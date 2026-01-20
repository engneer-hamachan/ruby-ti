package method_evaluator

import "ti/base"

var dynamicStrategies = make(map[[2]string]MethodEvaluateStrategy)

type MethodEvaluateStrategy interface {
	evaluate(m *MethodEvaluator) error
}

func NewStrategy(m *MethodEvaluator) MethodEvaluateStrategy {
	dynamicStrategy, ok :=
		dynamicStrategies[[2]string{
			m.evaluatedObjectT.GetObjectClass(),
			m.method,
		}]

	if ok {
		return dynamicStrategy
	}

	dynamicStrategy, ok =
		dynamicStrategies[[2]string{
			"",
			m.method,
		}]

	if ok {
		return dynamicStrategy
	}

	dynamicStrategy, ok =
		dynamicStrategies[[2]string{
			"Kernel",
			m.method,
		}]

	if ok {
		return dynamicStrategy
	}

	if m.objectT.ToString() == "union" {
		return &unionInstanceStrategy{}
	}

	if m.objectT.IsEmpty() {
		return &topLevelMethodStrategy{}
	}

	if m.evaluatedObjectT.IsClassType() {
		m.objectT = m.evaluatedObjectT.DeepCopy()
		return &classMethodStrategy{}
	}

	if m.objectT.IsClassType() {
		return &classMethodStrategy{}
	}

	// for single char class (e.g: class H end;)
	if base.IsClassDefined([]string{m.ctx.GetFrame()}, m.objectT.ToString()) && len(m.objectT.ToString()) == 1 {
		return &classMethodStrategy{}
	}

	if m.evaluatedObjectT.IsUnionType() {
		return &unionInstanceStrategy{}
	}

	return &instanceMethodStrategy{}
}
