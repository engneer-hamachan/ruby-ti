package base

import (
	"slices"
	"strings"
	"unicode"
)

func (t *T) IsEmptyDefineArgs() bool {
	return len(t.defineArgs) == 0
}

func (t *T) IsAcceptCountArgs(args []*T) bool {
	return len(t.defineArgs) == len(args)
}

func (t *T) IsTargetIdentifier(target string) bool {
	return t != nil && t.tType == UNKNOWN && t.ToString() == target
}

func (t *T) IsTargetIdentifiers(targets []string) bool {
	if t == nil || t.tType != UNKNOWN {
		return false
	}

	return slices.Contains(targets, t.ToString())
}

func (t *T) IsIdentifierType() bool {
	if t == nil {
		return false
	}

	return t.tType == UNKNOWN
}

func (t *T) IsUnknownType() bool {
	if t == nil {
		return false
	}

	return t.tType == UNKNOWN
}

func (t *T) IsClassType() bool {
	return t.tType == CLASS
}

func (t *T) IsClassIdentifier() bool {
	if t.tType != UNKNOWN {
		return false
	}

	if slices.Contains(BuiltinClasses, t.ToString()) {
		return true
	}

	var isClass bool
	if unicode.IsUpper(rune(t.ToString()[0])) {
		isClass = true
	}

	if isClass {
		for _, char := range t.ToString() {
			if unicode.IsLower(char) {
				return true
			}
		}
	}

	return false
}

func (t *T) IsConstType() bool {
	if t == nil {
		return false
	}

	return t.tType == CONST
}

func (t *T) IsConstIdentifier() bool {
	if t.tType != UNKNOWN {
		return false
	}

	if len(t.ToString()) < 1 {
		return false
	}

	if slices.Contains(BuiltinClasses, t.ToString()) {
		return false
	}

	var isConst bool
	if unicode.IsUpper(rune(t.ToString()[0])) {
		isConst = true
	}

	if len(t.ToString()) < 2 {
		return false
	}

	if isConst {
		for _, char := range t.ToString() {
			if char == ':' {
				return false
			}

			if unicode.IsLower(char) {
				return false
			}
		}
	}

	return isConst
}

func (t *T) IsDotIdentifier() bool {
	return t.IsTargetIdentifier(".")
}

func (t *T) IsAndDotIdentifier() bool {
	return t.IsTargetIdentifier("&.")
}

func (t *T) IsEndIdentifier() bool {
	return t.IsTargetIdentifier("end")
}

func (t *T) IsNewLineIdentifier() bool {
	return t.IsTargetIdentifier("\n")
}

func (t *T) IsEqualIdentifier() bool {
	return t.IsTargetIdentifier("=")
}

func (t *T) IsBoolIdentifier() bool {
	return t.IsTargetIdentifier("true") || t.IsTargetIdentifier("false")
}

func (t *T) IsCalcIdentifier() bool {
	calcIdentifiers := []string{"+", "-", "/", "*"}

	return t.IsTargetIdentifiers(calcIdentifiers)
}

func (t *T) IsCalcMethod() bool {
	calcIdentifiers := []string{"+", "-", "/", "*"}

	return slices.Contains(calcIdentifiers, t.GetMethodName())
}

func (t *T) IsRefferenceSquareT() bool {
	return t.IsTargetIdentifier("[") && !t.IsBeforeSpace
}

func (t *T) IsExclamationIdentifier() bool {
	return t.IsTargetIdentifier("!")
}

func (t *T) IsTransformTargetIdentifier() bool {
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

	return t.IsTargetIdentifiers(transformTargetIdentifiers)
}

func (t *T) IsTopLevelFunctionIdentifier(frame string, class string) bool {
	if t.tType != UNKNOWN {
		return false
	}

	// for hoge.class special case
	if t.IsTargetIdentifier("class") {
		return false
	}

	methodT := GetTopLevelMethodT(frame, class, t.ToString())

	if methodT == nil {
		methodT = GetTopLevelMethodT("", class, t.ToString())
	}

	return methodT != nil
}

func (t *T) IsCommaIdentifier() bool {
	return t.IsTargetIdentifier(",")
}

func (t *T) IsOpenParentheses() bool {
	return t.IsTargetIdentifier("(")
}

func (t *T) IsCloseParentheses() bool {
	return t.IsTargetIdentifier(")")
}

func (t *T) IsImmediate() bool {
	return t.tType != UNKNOWN
}

func (t *T) IsPredicateIdentifier() bool {
	predicateIdentifiers := []string{
		">",
		"<",
		"==",
		"!=",
		"<=",
		">=",
	}

	return t.IsTargetIdentifiers(predicateIdentifiers)
}

func (t *T) IsEmpty() bool {
	return t.ToString() == ""
}

func (t *T) IsBlockType() bool {
	if t == nil {
		return false
	}

	return t.tType == BLOCK
}

func (t *T) IsStringType() bool {
	if t == nil {
		return false
	}

	return t.tType == STRING
}

func (t *T) IsArrayType() bool {
	if t == nil {
		return false
	}

	return t.tType == ARRAY
}

func (t *T) IsHashType() bool {
	if t == nil {
		return false
	}

	return t.tType == HASH
}

func (t *T) IsRangeType() bool {
	if t == nil {
		return false
	}

	return t.tType == RANGE
}

func (t *T) IsIteratableType() bool {
	if t == nil {
		return false
	}

	iteratableTypes := []int{ARRAY, HASH, RANGE}

	return slices.Contains(iteratableTypes, t.tType)
}

func (t *T) IsUnionType() bool {
	if t == nil {
		return false
	}

	return t.tType == UNION
}

func (t *T) IsKeyValueType() bool {
	if t == nil {
		return false
	}

	return t.tType == KEYVALUE
}

func (t *T) IsMatchType(targetT *T) bool {
	return t.tType == targetT.tType

}

func (t *T) IsMatchUnionType(targetT *T) bool {
	switch targetT.tType {
	case UNION:
		targetTypes := targetT.GetUnionTypes()

		for _, candidateType := range t.GetUnionTypes() {
			if candidateType == UNTYPED {
				return true
			}

			if slices.Contains(targetTypes, candidateType) {
				continue
			}

			return false
		}

		return true

	default:
		for _, variantT := range t.variants {
			if variantT.IsAnyType() {
				return true
			}

			if variantT.tType == targetT.tType {
				return true
			}
		}

		return false
	}
}

func (t *T) HasDefault() bool {
	if t == nil {
		return false
	}

	return t.hasDefault
}

func (t *T) IsWhenCallType() bool {
	if t == nil {
		return false
	}

	return t.isWhenCallType
}

func (t *T) IsBuiltin() bool {
	if t == nil {
		return false
	}

	return t.isBuiltin
}

func (t *T) IsKeyIdentifier() bool {
	if t == nil {
		return false
	}

	if !t.IsIdentifierType() {
		return false
	}

	if len(t.ToString()) < 2 {
		return false
	}

	if t.ToString()[len(t.ToString())-1:] == ":" {
		return true
	}

	return false
}

func (t *T) IsSymbolIdentifier() bool {
	if t == nil {
		return false
	}

	if !t.IsIdentifierType() {
		return false
	}

	if len(t.ToString()) < 2 {
		return false
	}

	if t.ToString()[0] == ':' {
		return true
	}

	return false
}

func (t *T) IsReadOnly() bool {
	return t.isReadOnly
}

func (t *T) IsEqualObject(targetT *T) bool {
	targetObject := targetT.objectClass
	targetType := targetT.tType

	if len(t.variants) == 0 && len(targetT.variants) == 0 {
		return t.tType == targetType
	}

	for _, variantT := range t.variants {
		variantType := variantT.tType
		variantObject := variantT.objectClass

		if variantType == targetType && variantObject == targetObject {
			return true
		}
	}

	return false
}

func (t *T) IsTargetClassObject(target string) bool {
	if t == nil {
		return false
	}

	return t.objectClass == target
}

func (t *T) IsSymbolType() bool {
	if t == nil {
		return false
	}

	return t.tType == SYMBOL
}

func (t *T) IsAnyType() bool {
	if t == nil {
		return false
	}

	return t.tType == UNTYPED
}

func (t *T) IsAsteriskPrefix() bool {
	if t == nil {
		return false
	}

	if !t.IsIdentifierType() {
		return false
	}

	if len(t.ToString()) < 2 {
		return false
	}

	return t.ToString()[0] == '*'
}

func (t *T) IsBeforeEvaluateAsteriskPrefix() bool {
	if t == nil {
		return false
	}

	if len(t.beforeEvaluateCode) < 1 {
		return false
	}

	return t.beforeEvaluateCode[0] == '*'
}

func (t *T) IsBeforeEvaluateAtmarkPrefix() bool {
	if t == nil {
		return false
	}

	if len(t.beforeEvaluateCode) < 1 {
		return false
	}

	return t.beforeEvaluateCode[0] == '@'
}

func (t *T) IsNameSpaceIdentifier() bool {
	if t == nil {
		return false
	}

	if t.tType == STRING {
		return false
	}

	return len(strings.Split(t.ToString(), "::")) > 1
}

func (t *T) IsPriorityT() bool {
	return t.IsCalcIdentifier() || t.IsDotIdentifier() || t.IsAndDotIdentifier()
}

func (t *T) IsRefferenceAbleT() bool {
	return !t.IsTargetIdentifier("[") && !t.IsTargetIdentifier("\n")
}
