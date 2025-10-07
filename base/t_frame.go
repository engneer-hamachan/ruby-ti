package base

import (
	"maps"
)

var TFrame = make(map[FrameKey]*T)

type Sig struct {
	Contents string
	Detail   string
}

var TSignatures = []Sig{}

func DeepCopyTFrame() map[FrameKey]*T {
	copied := make(map[FrameKey]*T)
	maps.Copy(copied, TFrame)

	return copied
}

func RestoreFrame(currentFrame map[FrameKey]*T, originalFrame map[FrameKey]*T) {
	for key := range currentFrame {
		if _, ok := originalFrame[key]; !ok {
			delete(currentFrame, key)
		}
	}
}

func SetClassMethodT(
	frame string,
	class string,
	methodT *T,
	isPrivate bool,
) {

	info := frame + "::" + class + "." + methodT.method

	if len(methodT.GetDefineArgs()) == 0 {
		info += "()"
	} else {
		var args string
		args += "("
		for _, mt := range methodT.GetDefineArgs() {
			mmt := GetValueT(frame, class, methodT.GetMethodName(), mt)

			if args != "(" {
				args += " ,"
			}

			args += TypeToString(mmt)
		}
		args += ")"
		info += args
	}

	info += " -> "
	info += TypeToString(methodT)

	TSignatures = append(TSignatures, Sig{methodT.GetMethodName(), info})

	TFrame[classMethodTFrameKey(
		frame,
		class,
		methodT.GetMethodName(),
		isPrivate,
	)] = methodT
}

func SetMethodT(
	frame string,
	targetClass string,
	methodT *T,
	isPrivate bool,
) {

	switch targetClass {
	case "":
		methodT.SetBeforeEvaluateCode(methodT.GetMethodName())

	default:
		methodT.SetBeforeEvaluateCode(targetClass + "." + methodT.GetMethodName())
	}

	info := frame + "::" + targetClass + "." + methodT.method
	if len(methodT.GetDefineArgs()) == 0 {
		info += "()"
	} else {
		var args string
		args += "("
		for _, mt := range methodT.GetDefineArgs() {
			mmt := GetValueT(frame, targetClass, methodT.GetMethodName(), mt)

			if args != "(" {
				args += " ,"
			}

			args += TypeToString(mmt)
		}
		args += ")"
		info += args
	}

	info += " -> "
	info += TypeToString(methodT)

	TSignatures = append(TSignatures, Sig{methodT.GetMethodName(), info})

	TFrame[methodTFrameKey(
		frame,
		targetClass,
		methodT.GetMethodName(),
		isPrivate,
	)] = methodT
}

func getParentMethodT(
	frame string,
	class string,
	method string,
	isPrivate bool,
	isStatic bool,
) *T {

	classNode := ClassNode{frame, class}

	for _, parentNode := range ClassInheritanceMap[classNode] {
		var methodT *T
		var ok bool

		switch isStatic {
		case true:
			methodT, ok =
				TFrame[classMethodTFrameKey(
					parentNode.Frame,
					parentNode.Class,
					method,
					isPrivate,
				)]

		default:
			methodT, ok =
				TFrame[methodTFrameKey(
					parentNode.Frame,
					parentNode.Class,
					method,
					isPrivate,
				)]
		}

		if ok {
			return methodT
		}

		methodT =
			getParentMethodT(
				parentNode.Frame,
				parentNode.Class,
				method,
				isPrivate,
				isStatic,
			)

		if methodT != nil {
			return methodT
		}
	}

	return nil
}

func GetMethodT(frame, targetClass, targetMethod string, isPrivate bool) *T {
	methodT, ok :=
		TFrame[methodTFrameKey(frame, targetClass, targetMethod, isPrivate)]

	if ok {
		return methodT
	}

	methodT = getParentMethodT(frame, targetClass, targetMethod, isPrivate, false)
	if methodT != nil {
		return methodT
	}

	methodT, ok =
		TFrame[methodTFrameKey("Builtin", targetClass, targetMethod, isPrivate)]

	if ok {
		return methodT
	}

	methodT =
		getParentMethodT("Builtin", targetClass, targetMethod, isPrivate, false)

	if methodT != nil {
		return methodT
	}

	return methodT
}

func GetTopLevelMethodT(
	frame string,
	class, method string,
) *T {

	methodT, ok :=
		TFrame[methodTFrameKey(frame, class, method, false)]

	if ok {
		return methodT
	}

	methodT, ok =
		TFrame[classMethodTFrameKey(frame, class, method, false)]
	if ok {
		return methodT
	}

	methodT, ok =
		TFrame[methodTFrameKey(frame, class, method, true)]

	if ok {
		return methodT
	}

	methodT, ok =
		TFrame[classMethodTFrameKey(frame, class, method, true)]

	if ok {
		return methodT
	}

	methodT = getParentMethodT(frame, class, method, false, false)
	if methodT != nil {
		return methodT
	}

	if ok {
		return methodT
	}

	methodT, ok =
		TFrame[methodTFrameKey("Builtin", "", method, false)]

	if ok {
		return methodT
	}

	methodT = getParentMethodT("Builtin", "", method, false, false)
	if methodT != nil {
		return methodT
	}

	if ok {
		return methodT
	}

	return nil
}

func GetClassMethodT(
	frame string,
	targetClass string,
	targetMethod string,
	isPrivate bool,
) *T {

	methodT :=
		TFrame[classMethodTFrameKey(
			frame,
			targetClass,
			targetMethod,
			isPrivate,
		)]

	if methodT == nil {
		methodT =
			TFrame[classMethodTFrameKey(
				"Builtin",
				targetClass,
				targetMethod,
				isPrivate,
			)]
	}

	if methodT == nil {
		methodT =
			getParentMethodT(frame, targetClass, targetMethod, isPrivate, true)

		if methodT != nil {
			return methodT
		}
	}

	return methodT
}

func setParentValueT(
	frame string,
	class string,
	method string,
	variable string,
	t *T,
) bool {

	classNode := ClassNode{frame, class}
	for _, parentNode := range ClassInheritanceMap[classNode] {
		_, ok :=
			TFrame[valueTFrameKey(
				parentNode.Frame,
				parentNode.Class,
				method,
				variable,
			)]

		if ok {
			TFrame[valueTFrameKey(
				parentNode.Frame,
				parentNode.Class,
				method,
				variable,
			)] = t

			return true
		}

		ok =
			setParentValueT(parentNode.Frame, parentNode.Class, method, variable, t)

		if ok {
			return true
		}
	}

	return false
}

func SetValueT(
	frame string,
	class string,
	method string,
	variable string,
	t *T,
) error {

	if len(variable) > 0 && variable[0] == '*' {
		variable = variable[1:]
	}

	_, ok := TFrame[valueTFrameKey(frame, class, method, variable)]
	if ok {
		TFrame[valueTFrameKey(frame, class, method, variable)] = t

		return nil
	}

	ok = setParentValueT(frame, class, method, variable, t)
	if ok {
		return nil
	}

	TFrame[valueTFrameKey(frame, class, method, variable)] = t

	return nil
}

func getParentValueT(
	frame string,
	class string,
	method string,
	variable string,
) *T {

	classNode := ClassNode{frame, class}

	for _, parentNode := range ClassInheritanceMap[classNode] {
		t, ok :=
			TFrame[valueTFrameKey(
				parentNode.Frame,
				parentNode.Class,
				method,
				variable,
			)]

		if ok {
			return t
		}

		valueT :=
			getParentValueT(parentNode.Frame, parentNode.Class, method, variable)

		if valueT != nil {
			return valueT
		}
	}

	return nil
}

func GetValueT(frame string, class string, method string, variable string) *T {
	if len(variable) > 0 && variable[0] == '*' {
		variable = variable[1:]
	}

	t := TFrame[valueTFrameKey(frame, class, method, variable)]

	if t != nil {
		return t
	}

	t = getParentValueT(frame, class, method, variable)

	if t != nil {
		return t
	}

	return getParentValueT("Builtin", class, method, variable)
}

func SetInstanceValueT(
	frame string,
	class string,
	variable string,
	t *T,
) {

	TFrame[valueTFrameKey(frame, class, "", variable)] = t
}

func GetInstanceValueT(frame string, class string, variable string) *T {
	t := TFrame[valueTFrameKey(frame, class, "", variable)]
	if t != nil {
		return t
	}

	return getParentValueT(frame, class, "", variable)
}

func CalculateFrame(frame string, class string) string {
	switch {
	case frame == "" && class == "":
		return ""

	case frame == "" && class != "":
		return class

	case frame != "" && class == "":
		return frame

	default:
		return frame + "::" + class
	}
}

func SetConstValueT(frame string, class string, variable string, t *T) {
	frame = CalculateFrame(frame, class)

	TFrame[constTFrameKey(frame, variable)] = t
}

func GetConstValueT(frame string, class string, variable string) *T {
	frame = CalculateFrame(frame, class)

	t, ok := TFrame[constTFrameKey(frame, variable)]
	if ok {
		return t
	}

	t, ok = TFrame[constTFrameKey("Builtin"+"::"+frame, variable)]
	if ok {
		return t
	}

	return nil
}

func GetDynamicValueT(
	frame string,
	class string,
	method string,
	instance string,
) *T {

	evaluatedObjectT :=
		GetValueT(frame, class, method, instance)

	if evaluatedObjectT == nil {
		evaluatedObjectT = GetTopLevelMethodT(frame, class, instance)
	}

	if evaluatedObjectT == nil && instance != "" && instance[0] == '@' {
		evaluatedObjectT = GetInstanceValueT(frame, class, instance[1:])
	}

	if evaluatedObjectT == nil {
		return MakeIdentifier(instance)
	}

	return evaluatedObjectT
}
