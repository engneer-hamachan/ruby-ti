package base

type FrameKey struct {
	frame          string
	targetClass    string
	targetMethod   string
	targetVariable string
	isPrivate      bool
	isStatic       bool
}

type PubFrameKey struct {
	Frame          string
	TargetClass    string
	TargetMethod   string
	TargetVariable string
	IsStatic       bool
}

func (f *FrameKey) Variable() string {
	return f.targetVariable
}

func classMethodTFrameKey(
	frame string,
	targetClass string,
	targetMethod string,
	isPrivate bool,
) FrameKey {

	return FrameKey{
		frame:        frame,
		targetClass:  targetClass,
		targetMethod: targetMethod,
		isPrivate:    isPrivate,
		isStatic:     true,
	}
}

func methodTFrameKey(
	frame string,
	targetClass string,
	targetMethod string,
	isPrivate bool,
) FrameKey {

	return FrameKey{
		frame:        frame,
		targetClass:  targetClass,
		targetMethod: targetMethod,
		isPrivate:    isPrivate,
	}
}

func valueTFrameKey(
	frame string,
	targetClass string,
	targetMethod string,
	targetVariable string,
	isStatic bool,
) FrameKey {

	return FrameKey{
		frame:          frame,
		targetClass:    targetClass,
		targetMethod:   targetMethod,
		targetVariable: targetVariable,
		isStatic:       isStatic,
	}
}

func constTFrameKey(
	frame string,
	targetVariable string,
) FrameKey {

	return FrameKey{
		frame:          frame,
		targetVariable: targetVariable,
	}
}
