package base

type T struct {
	ID                  string
	tType               int
	val                 any
	key                 string
	frame               string
	objectClass         string
	method              string
	defineArgs          []string
	hasDefault          bool
	isBuiltin           bool
	variants            []T
	isWhenCallType      bool
	IsBuiltinAsterisk   bool
	IsConditionalReturn bool
	IsDestructive       bool
	beforeEvaluateCode  string
	isReadOnly          bool
	blockParamaters     []T
	IsBeforeSpace       bool
	IsBlockGiven        bool
	IsProtected         bool
	IsStatic            bool
	DefinedFrame        string
	DefinedClass        string
	DefinedMethod       string
	IsCaptureOwner      bool
	owner               any
}

func (t *T) DeepCopy() *T {
	if t == nil {
		return nil
	}

	result := &T{
		tType:              t.tType,
		val:                t.val,
		key:                t.key,
		frame:              t.frame,
		objectClass:        t.objectClass,
		method:             t.method,
		hasDefault:         t.hasDefault,
		isBuiltin:          t.isBuiltin,
		isWhenCallType:     t.isWhenCallType,
		beforeEvaluateCode: t.beforeEvaluateCode,
		isReadOnly:         t.isReadOnly,
		IsBeforeSpace:      t.IsBeforeSpace,
		IsBlockGiven:       t.IsBlockGiven,
		IsProtected:        t.IsProtected,
		DefinedFrame:       t.DefinedFrame,
		DefinedClass:       t.DefinedClass,
	}

	if t.defineArgs != nil {
		result.defineArgs = make([]string, len(t.defineArgs))
		copy(result.defineArgs, t.defineArgs)
	}

	if t.variants != nil {
		result.variants = make([]T, len(t.variants))
		for i, variant := range t.variants {
			result.variants[i] = *variant.DeepCopy()
		}
	}

	if t.blockParamaters != nil {
		result.blockParamaters = make([]T, len(t.blockParamaters))
		for i, param := range t.blockParamaters {
			result.blockParamaters[i] = *param.DeepCopy()
		}
	}

	return result
}
