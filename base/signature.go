package base

import (
	"slices"
)

type Sig struct {
	Method    string
	Detail    string
	Frame     string
	Class     string
	IsStatic  bool
	IsPrivate bool
	FileName  string
	Row       int
	Document  string
}

var TSignatures = make(map[string]Sig)
var TSignatureDocument = make(map[string]string)
var MethodCallPoint = make(map[string][]string)

func GetSortedTSignatures() []Sig {
	sortedSignatures := make([]Sig, 0, len(TSignatures))

	for _, sig := range TSignatures {
		sortedSignatures = append(sortedSignatures, sig)
	}

	slices.SortFunc(sortedSignatures, func(a, b Sig) int {
		if a.Method != b.Method {
			if a.Method < b.Method {
				return -1
			}
			return 1
		}
		if a.Class != b.Class {
			if a.Class < b.Class {
				return -1
			}
			return 1
		}
		if a.Frame < b.Frame {
			return -1
		}
		if a.Frame > b.Frame {
			return 1
		}
		return 0
	})

	return sortedSignatures
}

func TypeToStringForSignature(t *T) string {
	var content string

	if t.HasDefault() {
		content += "optional "
	}

	if t.IsAsteriskPrefix() {
		content += "*"
	}

	if t.IsDoubleAsteriskPrefix() {
		content += "**"
	}

	switch t.GetType() {
	case UNKNOWN:
		if t.IsUpperPrefix() {
			content += t.ToString()
			break
		}

		content += "unknown"

	default:
		content += TypeToString(t)
	}

	return content
}

func makeConditionalReturnContent(
	frame string,
	class string,
	methodT *T,
	content string,
) string {

	content += " "

	returnVariants := methodT.GetVariants()

	snapShot := content

	for _, darg := range methodT.GetDefineArgs() {
		if snapShot != content {
			content += ", "
		}

		dargCopy := darg

		if IsKeySuffix(dargCopy) {
			dargCopy = RemoveSuffix(dargCopy)
		}

		dargT :=
			GetValueT(frame, class, methodT.GetMethodName(), dargCopy, methodT.IsStatic)

		if dargT.IsUnionType() {
			for idx, variant := range dargT.GetVariants() {
				if idx > 0 {
					content += ", "
				}

				content += "("
				content += TypeToStringForSignature(&variant)
				content += ")"
				content += " -> "

				if idx < len(returnVariants) {
					content += TypeToStringForSignature(&returnVariants[idx])
				} else {
					content += TypeToStringForSignature(methodT)
				}
			}

			return content
		}
	}

	for i, returnVariant := range returnVariants {
		if snapShot != content {
			content += ", "
		}

		content += "("

		argIdx := 0
		if i < len(methodT.GetDefineArgs()) {
			argIdx = i
		}

		for j, darg := range methodT.GetDefineArgs() {
			if j > 0 {
				content += ", "
			}

			if j == argIdx {
				content += TypeToStringForSignature(&returnVariant)
			} else {
				dargCopy := darg

				if IsKeySuffix(dargCopy) {
					dargCopy = RemoveSuffix(dargCopy)
				}

				dargT :=
					GetValueT(
						frame,
						class,
						methodT.GetMethodName(),
						dargCopy,
						methodT.IsStatic,
					)

				content += TypeToStringForSignature(dargT)
			}
		}

		content += ") -> "
		content += TypeToStringForSignature(&returnVariant)
	}

	return content
}

func MakeSignatureContent(
	prefix string,
	frame, class string,
	methodT *T,
) string {

	content := prefix
	var args string

	switch methodT.IsConditionalReturn {
	case true:
		args += "(Match) =>"

	default:
		args += "("

		for _, darg := range methodT.GetDefineArgs() {
			if args != "(" {
				args += ", "
			}

			if IsKeySuffix(darg) {
				args += darg + " "
				darg = RemoveSuffix(darg)
			}

			dargT :=
				GetValueT(frame, class, methodT.GetMethodName(), darg, methodT.IsStatic)

			// *a or **a
			if darg[0] == '*' {
				switch darg[1] {
				case '*':
					dargT = MakeDoubleAsteriskKeyValue()

				default:
					dargT = MakeAsteriskUntyped()
				}
			}

			args += TypeToStringForSignature(dargT)
		}

		args += ")"
	}

	content += args

	if methodT.IsBlockGiven {
		content += " <block_params: "

		snapShot := content

		if len(methodT.GetBlockParameters()) == 0 {
			content += "void"
		}

		for _, variant := range methodT.GetBlockParameters() {
			if snapShot != content {
				content += ", "
			}

			content += TypeToStringForSignature(&variant)
		}

		content += ">"
	}

	if methodT.IsConditionalReturn {
		return makeConditionalReturnContent(frame, class, methodT, content)
	}

	content += " -> "
	content += TypeToStringForSignature(methodT)

	return content
}

func appendSignature(
	frame, class string,
	methodT *T,
	isStatic bool,
	isPrivate bool,
	fileName string,
	row int,
) {

	key := frame + class + methodT.method
	if isStatic {
		key += "static"
	}

	content := MakeSignatureContent(methodT.method, frame, class, methodT)
	document := TSignatureDocument[key]

	sig :=
		Sig{
			methodT.GetMethodName(),
			content,
			frame,
			class,
			isStatic,
			isPrivate,
			fileName,
			row,
			document,
		}

	TSignatures[key] = sig
}

func (s *Sig) GetPrintDetail() string {
	if s.Class != "" {
		return s.Class + "." + s.Detail
	}

	return s.Detail
}
