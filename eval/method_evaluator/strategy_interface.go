package method_evaluator

var dynamicStrategies = make(map[[2]string]MethodEvaluateStrategy)

type MethodEvaluateStrategy interface {
	evaluate(m *MethodEvaluator) error
}

func NewStrategy(m *MethodEvaluator) MethodEvaluateStrategy {
	// TODO: 継承から引くようにする
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
	//ここまで

	if m.objectT.ToString() == "union" {
		return &unionInstanceStrategy{}
	}

	if m.objectT.IsEmpty() {
		return &topLevelMethodStrategy{}
	}

	if m.objectT.IsClassType() {
		return &classMethodStrategy{}
	}

	if m.evaluatedObjectT.IsUnionType() {
		return &unionInstanceStrategy{}
	}

	return &instanceMethodStrategy{}
}
