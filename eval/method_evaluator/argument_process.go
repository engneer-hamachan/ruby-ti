package method_evaluator

import (
	"sort"
	"ti/base"
	"ti/context"
)

func (m *MethodEvaluator) isSkipT(t *base.T) bool {
	if !t.IsIdentifierType() {
		return false
	}

	id := t.ToString()

	switch id {
	case ",":
		return true

	case "\n":
		if m.isParentheses {
			return true
		}
	}

	return false
}

func (m *MethodEvaluator) isNotArgT(
	methodT *base.T,
	argTs []*base.T,
	t *base.T,
) bool {

	if methodT.IsCalcMethod() && len(argTs) >= 1 {
		m.parser.Unget()
		return true
	}

	if t.IsCommaIdentifier() || t.IsPredicateIdentifier() {
		if len(argTs) < 1 {
			return true
		}
	}

	if t.IsTargetIdentifier("[") && methodT.IsEmptyDefineArgs() {
		m.parser.Unget()
		return true
	}

	if !m.isParentheses && t.IsTargetIdentifier(")") {
		m.parser.Unget()
		return true
	}

	if t.IsTargetIdentifier("do") {
		m.parser.Unget()
		return true
	}

	if t.IsTargetIdentifier("{") && !m.isParentheses {
		m.parser.Unget()
		return true
	}

	if t.IsTargetIdentifier("if") {
		m.parser.Unget()
		return true
	}

	if t.IsTargetIdentifier("}") {
		m.parser.Unget()
		return true
	}

	if t.GetPower() > 0 && !t.IsTargetIdentifier("!") {
		m.parser.Unget()
		return true
	}

	return false
}

func (m *MethodEvaluator) isEndOfCollectArgs(t *base.T) bool {
	if t == nil {
		return true
	}

	if !t.IsIdentifierType() {
		return false
	}

	var endIdentifier string

	switch m.isParentheses {
	case true:
		endIdentifier = ")"

	default:
		endIdentifier = "\n"
	}

	return t.IsTargetIdentifier(endIdentifier)
}

func (m *MethodEvaluator) parseKeyIdentifierToKeyWordT(
	nextT *base.T,
) (*base.T, error) {

	t, err := m.parser.Read()
	if err != nil {
		return nil, err
	}

	// test(a: 1)
	if !t.IsCloseParentheses() && !t.IsCommaIdentifier() {
		err = m.outerEval.Eval(m.parser, m.ctx, t)
		if err != nil {
			return nil, err
		}

		valueT := m.parser.GetLastEvaluatedT()

		return base.MakeKeyValue(nextT.ToString(), &valueT), nil
	}

	// test(a:)
	m.parser.Unget()

	t = base.MakeIdentifier(nextT.ToRemoveSuffixString())

	err = m.outerEval.Eval(m.parser, m.ctx, t)
	if err != nil {
		return nil, err
	}

	valueT := m.parser.GetLastEvaluatedT()

	return base.MakeKeyValue(nextT.ToString(), &valueT), nil
}

func (m *MethodEvaluator) makeNextArg(nextT *base.T) (*base.T, error) {
	var err error

	for {
		if m.isSkipT(nextT) {
			nextT, err = m.parser.Read()
			if err != nil {
				return nil, err
			}

			continue
		}

		if nextT.IsKeyIdentifier() {
			nextT, err = m.parseKeyIdentifierToKeyWordT(nextT)
			if err != nil {
				return nil, err
			}
		}

		break
	}

	return nextT, nil
}

func splatArg(
	m *MethodEvaluator,
	ctx context.Context,
	nextT *base.T,
	argTs []*base.T,
) ([]*base.T, error) {

	err := m.outerEval.Eval(
		m.parser, ctx, base.MakeIdentifier(nextT.ToString()[1:]),
	)
	if err != nil {
		return argTs, err
	}

	lastEvaluatedT := m.parser.GetLastEvaluatedT()

	if !lastEvaluatedT.IsArrayType() {
		argTs = append(argTs, &lastEvaluatedT)
		return argTs, nil
	}

	unionT := lastEvaluatedT.UnifyVariants()

	for {
		if len(argTs) >= len(lastEvaluatedT.GetVariants()) {
			break
		}

		argTs = append(argTs, unionT)
	}

	return argTs, nil
}

func checkParentheses(m *MethodEvaluator) (bool, error) {
	nextT, err := m.parser.Read()
	if err != nil {
		return false, err
	}

	if nextT.IsOpenParentheses() && !nextT.IsBeforeSpace {
		return true, nil
	}

	m.parser.Unget()

	return false, nil
}

func getEvaluatedDefineArgs(
	m *MethodEvaluator,
	class string,
	methodT *base.T,
) []*base.T {

	var defineArgTs []*base.T

	for _, definedArg := range methodT.GetDefineArgs() {
		definedArgT :=
			base.GetValueT(
				methodT.GetFrame(),
				class,
				m.method,
				definedArg,
			)

		defineArgTs = append(defineArgTs, definedArgT)
	}

	return defineArgTs
}

func expectBlockArgProcess(
	m *MethodEvaluator,
	methodT *base.T,
	argTs []*base.T,
) ([]*base.T, error) {

	nextT, err := m.parser.Read()
	if err != nil {
		return argTs, err
	}

	if nextT.IsNewLineIdentifier() && !m.isParentheses {
		m.parser.Unget()
		return argTs, nil
	}

	if nextT.IsTargetIdentifier("do") || nextT.IsTargetIdentifier("{") {
		m.parser.SetLastEvaluatedT(m.evaluatedObjectT)

		m.parser.SetTmpEvaluaetdArgs(argTs)
		defer m.parser.ClearTmpEvaluaetdArgs()

		err = m.outerEval.Eval(m.parser, m.ctx, base.MakeIdentifier("do"))
		if err != nil {
			return argTs, err
		}

		return argTs, nil
	}

	if methodT.IsBlockGiven {
		err := m.outerEval.Eval(m.parser, m.ctx, nextT)
		if err != nil {
			return argTs, err
		}

		lastEvaluatedT := m.parser.GetLastEvaluatedT()

		err =
			makeTypeError(
				m.ctx.GetClass(),
				methodT.GetBeforeEvaluateCode(),
				base.TypeToString(&lastEvaluatedT),
				"Block",
			)

		return argTs, err
	}

	m.parser.Unget()

	return argTs, nil
}

func getEvaluatedArgs(
	m *MethodEvaluator,
	methodT *base.T,
) (argTs []*base.T, err error) {

	m.parser.EndParsingExpression()

	m.ctx.StartCallArg()
	defer m.ctx.EndCallArg()

	m.isParentheses, err = checkParentheses(m)
	if err != nil {
		return argTs, err
	}

	if methodT.IsEmptyDefineArgs() && !m.isParentheses {
		return expectBlockArgProcess(m, methodT, argTs)
	}

	for {
		t, err := m.parser.Read()
		if err != nil {
			return argTs, err
		}

		if m.isNotArgT(methodT, argTs, t) {
			break
		}

		nextT, err := m.makeNextArg(t)
		if err != nil {
			return argTs, err
		}

		if m.isEndOfCollectArgs(nextT) {
			if !m.isParentheses {
				m.parser.Unget()
			}

			break
		}

		if nextT.IsAsteriskPrefix() {
			argTs, err = splatArg(m, m.ctx, nextT, argTs)
			if err != nil {
				return argTs, err
			}

			continue
		}

		err = m.outerEval.Eval(m.parser, m.ctx, nextT)
		if err != nil {
			return argTs, err
		}

		lastEvaluatedT := m.parser.GetLastEvaluatedT()
		argTs = append(argTs, &lastEvaluatedT)
	}

	return expectBlockArgProcess(m, methodT, argTs)
}

func sortTsByKey(tList []*base.T) []*base.T {
	sort.Slice(tList, func(i, j int) bool {
		return tList[i].GetKey() < tList[j].GetKey()
	})

	return tList
}

func sortStrings(strList []string) []string {
	sort.Strings(strList)
	return strList
}

func prioritizeDefineArgNames(
	definedArgNames []string,
) []string {

	var namedDefineArgs []string
	var defineArgs []string

	for _, name := range definedArgNames {
		if name[len(name)-1:] == ":" && len(name) >= 2 {
			namedDefineArgs = append(namedDefineArgs, name)
			continue
		}

		defineArgs = append(defineArgs, name)
	}

	namedDefineArgs = sortStrings(namedDefineArgs)

	prioritizeDefineArgs :=
		append(defineArgs, namedDefineArgs...)

	return prioritizeDefineArgs
}

func prioritizeArgTs(
	argTs []*base.T,
) []*base.T {

	var namedArgTs []*base.T
	var otherArgTs []*base.T

	for _, t := range argTs {
		if t.IsKeyValueType() {
			namedArgTs = append(namedArgTs, t)
			continue
		}

		otherArgTs = append(otherArgTs, t)
	}

	namedArgTs = sortTsByKey(namedArgTs)

	prioritizeArgTs := append(otherArgTs, namedArgTs...)

	return prioritizeArgTs
}
