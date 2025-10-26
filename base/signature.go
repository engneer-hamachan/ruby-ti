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

func appendSignature(
	frame, class string,
	methodT *T,
	isStatic bool,
	fileName string,
	row int,
) {

	info := methodT.method
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
	info += args

	if methodT.IsBlockGiven {
		info += " <block_params: "

		snapShot := info

		for _, variant := range methodT.GetBlockParameters() {
			if snapShot != info {
				info += ", "
			}

			info += TypeToString(&variant)
		}

		info += ">"
	}

	info += " -> "
	info += TypeToString(methodT)

	sig :=
		Sig{methodT.GetMethodName(), info, frame, class, isStatic, fileName, row}

	key := frame + class + methodT.method

	if isStatic {
		key += "static"
	}

	TSignatures[key] = sig
}
