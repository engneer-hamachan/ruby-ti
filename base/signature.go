package base

import "slices"

type Sig struct {
	Method   string
	Detail   string
	Frame    string
	Class    string
	IsStatic bool
	FileName string
	Row      int
}

var TSignatures = make(map[string]Sig)

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

func MakeSignatureContent(
	prefix string,
	frame, class string,
	methodT *T,
) string {

	content := prefix
	var args string

	args += "("
	for _, darg := range methodT.GetDefineArgs() {
		if args != "(" {
			args += ", "
		}

		if isKeySuffix(darg) {
			args += darg + " "
			darg = removeSuffix(darg)
		}

		dargT :=
			GetValueT(frame, class, methodT.GetMethodName(), darg, methodT.IsStatic)

		if dargT.HasDefault() {
			args += `optional `
		}

		if dargT.IsAsteriskPrefix() {
			args += "*"
		}

		switch dargT.GetType() {
		case UNION:
			args += UnionTypeToString(dargT.GetVariants())

		case UNKNOWN:
			args += `unknown`

		default:
			args += TypeToString(dargT)
		}
	}

	args += ")"
	content += args

	if methodT.IsBlockGiven {
		content += " <block_params: "

		snapShot := content

		for _, variant := range methodT.GetBlockParameters() {
			if snapShot != content {
				content += ", "
			}

			content += TypeToString(&variant)
		}

		content += ">"
	}

	content += " -> "

	switch methodT.GetType() {
	case UNION:
		content += UnionTypeToString(methodT.GetVariants())

	case UNKNOWN:
		content += "?"

	default:
		content += TypeToString(methodT)
	}

	return content
}

func appendSignature(
	frame, class string,
	methodT *T,
	isStatic bool,
	fileName string,
	row int,
) {

	content := MakeSignatureContent(methodT.method, frame, class, methodT)

	sig :=
		Sig{methodT.GetMethodName(), content, frame, class, isStatic, fileName, row}

	key := frame + class + methodT.method

	if isStatic {
		key += "static"
	}

	TSignatures[key] = sig
}
