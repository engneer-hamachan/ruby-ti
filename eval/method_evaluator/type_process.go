package method_evaluator

import (
	"fmt"
	"ti/base"
)

func isKeySuffix(str string) bool {
	return str[len(str)-1:] == ":" && len(str) >= 2
}

func isAsteriskPrefix(str string) bool {
	return str[0] == '*'
}

func isAcceptIdx(tSlice []*base.T, idx int) bool {
	return len(tSlice) >= (idx + 1)
}

func isNotAcceptIdx(tSlice []*base.T, idx int) bool {
	return len(tSlice) < (idx + 1)
}

func isKeyValueTypeArgBeforeAcceptIdx(argTs []*base.T, argIdx int) bool {
	if !isAcceptIdx(argTs, argIdx) {
		return false
	}

	return argTs[argIdx].IsKeyValueType()
}

func isExtraArgError(
	isKeyTypeDefineArg bool,
	argTs []*base.T,
	argIdx int,
) bool {

	if !isKeyTypeDefineArg {
		return false
	}

	if !isAcceptIdx(argTs, argIdx) {
		return false
	}

	if !argTs[argIdx].IsKeyValueType() {
		return true
	}

	return false
}

func isNotDefineArgArgsError(
	isKeyTypeDefineArg bool,
	argTs []*base.T,
	argIdx int,
) bool {

	return isKeyValueTypeArgBeforeAcceptIdx(argTs, argIdx) && !isKeyTypeDefineArg
}

func isNotContainDefineArg(argTs []*base.T, definedArg string) bool {
	var isContain bool

	for _, argT := range argTs {
		if argT.IsKeyValueType() && argT.GetRemoveSuffixKey() == definedArg {
			isContain = true
		}
	}

	return !isContain
}

func isNotDefineNamedArgError(
	isKeyTypeDefineArg bool,
	definedArgT *base.T,
	argTs []*base.T,
	definedArg string,
) bool {

	if !isKeyTypeDefineArg {
		return false
	}

	if definedArgT.HasDefault() {
		return false
	}

	if !isNotContainDefineArg(argTs, definedArg) {
		return false
	}

	return true
}

func removeSuffix(str string) string {
	return str[:len(str)-1]
}

func propagationForCalledTo(
	m *MethodEvaluator,
	class, definedArg string,
	definedArgT *base.T,
	argT *base.T,
) bool {

	if argT.IsIdentifierType() {
		return false
	}

	if definedArgT == nil || definedArgT.IsIdentifierType() {
		argT.SetIsWhenCallType(true)

		base.SetValueT(
			m.evaluatedObjectT.GetFrame(),
			class,
			m.method,
			definedArg,
			argT,
		)

		return true
	}

	if definedArgT.IsBuiltin() {
		return false
	}

	if definedArgT.IsUnionType() && definedArgT.HasDefault() {
		definedArgT.AppendVariant(*argT)

		return true
	}

	if definedArgT.IsUnionType() && definedArgT.IsWhenCallType() {
		definedArgT.AppendVariant(*argT)

		return true
	}

	if definedArgT.IsMatchType(argT) {
		return true
	}

	if definedArgT.HasDefault() || definedArgT.IsWhenCallType() {
		var unionVariants []base.T

		switch definedArgT.GetType() {
		case base.UNION:
			unionVariants = append(unionVariants, definedArgT.GetVariants()...)
		default:
			unionVariants = append(unionVariants, *definedArgT)
		}

		switch argT.GetType() {
		case base.UNION:
			unionVariants = append(unionVariants, argT.GetVariants()...)
		default:
			unionVariants = append(unionVariants, *argT)
		}

		unionT := base.MakeUnion(unionVariants)

		unionT.SetHasDefault(definedArgT.HasDefault())
		unionT.SetIsWhenCallType(definedArgT.IsWhenCallType())

		base.SetValueT(
			m.evaluatedObjectT.GetFrame(),
			class,
			m.method,
			definedArg,
			unionT,
		)

		return true
	}

	return false
}

func propagationForCaller(
	m *MethodEvaluator,
	definedArgT *base.T,
	argT *base.T,
) bool {

	if definedArgT.IsAnyType() {
		return true
	}

	if !argT.IsIdentifierType() || argT.IsConstIdentifier() {
		return false
	}

	base.SetValueT(
		m.ctx.GetFrame(),
		m.ctx.GetClass(),
		m.ctx.GetMethod(),
		argT.ToString(),
		definedArgT,
	)

	return true
}

func checkArgType(
	m *MethodEvaluator,
	class string,
	definedArgT *base.T,
	argT *base.T,
) error {

	switch {
	case argT.IsBlockType():
		return nil

	case definedArgT.IsAnyType() || argT.IsAnyType():
		return nil

	case definedArgT.IsMatchType(argT):
		return nil

	case definedArgT.IsUnionType():
		if definedArgT.IsMatchUnionType(argT) {
			return nil
		}

		var argType string

		switch argT.GetType() {
		case base.UNION:
			argType = base.UnionTypeToString(argT.GetVariants())

		default:
			argType = base.TypeToString(argT)
		}

		return makeTypeError(
			class,
			m.method,
			argType,
			base.UnionTypeToString(definedArgT.GetVariants()),
		)

	case argT.IsUnionType():
		if argT.IsMatchUnionType(definedArgT) {
			return nil
		}

		return makeTypeError(
			class,
			m.method,
			base.UnionTypeToString(argT.GetVariants()),
			base.TypeToString(definedArgT),
		)

	default:
		return makeTypeError(
			class,
			m.method,
			base.TypeToString(argT),
			base.TypeToString(definedArgT),
		)
	}
}

func asteriskDefineProcess(
	m *MethodEvaluator,
	class string,
	definedArgNames []string,
	defineArgIdx int,
	argTs []*base.T,
	argIdx int,
) (int, int) {

	asteriskArrayT := base.MakeAnyArray()
	mustBindCt := len(definedArgNames[defineArgIdx+1:])
	remainingArgTs := argTs[argIdx:]

	if mustBindCt >= len(remainingArgTs) {
		base.SetValueT(
			m.evaluatedObjectT.GetFrame(),
			class,
			m.method,
			definedArgNames[defineArgIdx][1:],
			asteriskArrayT,
		)

		defineArgIdx++

		return defineArgIdx, argIdx
	}

	for idx, argT := range remainingArgTs {
		asteriskArrayT.AppendArrayVariant(*argT)

		if mustBindCt == len(remainingArgTs[idx+1:]) {
			break
		}

		argIdx++
	}

	base.SetValueT(
		m.evaluatedObjectT.GetFrame(),
		class,
		m.method,
		definedArgNames[defineArgIdx][1:],
		asteriskArrayT,
	)

	argIdx++
	defineArgIdx++

	return defineArgIdx, argIdx
}

func checkAndPropagateArgs(
	m *MethodEvaluator,
	class string,
	methodT *base.T,
	argTs []*base.T,
) (err error) {

	var isAsterisk bool
	var defineArgIdx int
	var argIdx int

	sortedDfineArgs := prioritizeDefineArgNames(methodT.GetDefineArgs())
	sortedArgTs := prioritizeArgTs(argTs)

	for {
		if (defineArgIdx + 1) > len(sortedDfineArgs) {
			break
		}

		definedArg := sortedDfineArgs[defineArgIdx]

		// *a
		if isAsteriskPrefix(definedArg) {
			if len(sortedArgTs) < argIdx {
				break
			}

			isAsterisk = true

			defineArgIdx, argIdx =
				asteriskDefineProcess(
					m,
					class,
					sortedDfineArgs,
					defineArgIdx,
					sortedArgTs,
					argIdx,
				)

			continue
		}

		//a:
		isKeyTypeDefineArg := isKeySuffix(definedArg)

		if isKeyTypeDefineArg {
			definedArg = removeSuffix(definedArg)
		}

		definedArgT :=
			base.GetValueT(
				methodT.GetFrame(),
				class,
				m.method,
				definedArg,
			)

		// TODO: design later
		if definedArgT == nil {
			definedArgT =
				base.GetValueT(
					methodT.DefinedFrame,
					methodT.DefinedClass,
					m.method,
					definedArg,
				)
		}

		if isNotDefineNamedArgError(
			isKeyTypeDefineArg,
			definedArgT,
			sortedArgTs,
			definedArg,
		) {

			err = fmt.Errorf(
				"%s: is not defined expected %s",
				definedArg,
				makeDefineArgumentInfo(m, class, methodT),
			)
			m.parser.Fatal(m.ctx, err)
		}

		if isExtraArgError(isKeyTypeDefineArg, sortedArgTs, argIdx) {
			err =
				fmt.Errorf(
					"%s is extra argument",
					sortedArgTs[argIdx].GetBeforeEvaluateCode(),
				)

			m.parser.Fatal(m.ctx, err)
		}

		if isNotDefineArgArgsError(isKeyTypeDefineArg, sortedArgTs, argIdx) {
			err = fmt.Errorf(
				"%s is not defined expected %s",
				definedArg,
				makeDefineArgumentInfo(m, class, methodT),
			)
			m.parser.Fatal(m.ctx, err)
		}

		if isKeyValueTypeArgBeforeAcceptIdx(sortedArgTs, argIdx) {
			switch sortedArgTs[argIdx].GetRemoveSuffixKey() {
			case definedArg:
				sortedArgTs[argIdx] = sortedArgTs[argIdx].GetKeyValue()

			default:
				if definedArgT.HasDefault() {
					defineArgIdx++
					continue
				}
			}
		}

		if isNotAcceptIdx(sortedArgTs, argIdx) {
			switch definedArgT.HasDefault() {
			case true:
				argIdx++
				defineArgIdx++

				continue

			default:
				err = fmt.Errorf(
					"too few arguments for %s expected %s",
					methodT.GetBeforeEvaluateCode(),
					makeDefineArgumentInfo(m, class, methodT),
				)

				m.parser.Fatal(m.ctx, err)
			}
		}

		if len(sortedArgTs) <= argIdx {
			break
		}

		if propagationForCalledTo(
			m,
			class,
			definedArg,
			definedArgT,
			sortedArgTs[argIdx],
		) {

			argIdx++
			defineArgIdx++

			continue
		}

		if propagationForCaller(m, definedArgT, sortedArgTs[argIdx]) {
			argIdx++
			defineArgIdx++

			continue
		}

		err := checkArgType(m, class, definedArgT, sortedArgTs[argIdx])
		if err != nil {
			return err
		}

		argIdx++
		defineArgIdx++
	}

	if len(sortedArgTs) > len(methodT.GetDefineArgs()) && !isAsterisk {
		err =
			fmt.Errorf("too many arguments for %s", methodT.GetBeforeEvaluateCode())

		m.parser.Fatal(m.ctx, err)
	}

	return nil
}

func checkAndPropagateArgsForUnionWithReturnT(
	m *MethodEvaluator,
	classNames []string,
	methodTs []*base.T,
	evalutedArgs []*base.T,
) (returnT *base.T, err error) {

	for idx, class := range classNames {
		err = checkAndPropagateArgs(m, class, methodTs[idx], evalutedArgs)
		if err != nil {
			return nil, err
		}

		if returnT == nil {
			returnT = methodTs[idx]

			continue
		}

		if returnT.IsUnionType() {
			returnT.AppendVariant(*methodTs[idx])

			continue
		}

		if methodTs[idx].IsUnionType() {
			methodTs[idx].AppendVariant(*returnT)

			returnT = base.MakeUnion(methodTs[idx].GetVariants())

			continue
		}

		if !returnT.IsMatchType(methodTs[idx]) {
			returnT = base.MakeUnion([]base.T{*returnT, *methodTs[idx]})

			continue
		}
	}

	return returnT, nil
}
