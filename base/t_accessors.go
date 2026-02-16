package base

// Basic type and field accessors
func (t *T) GetType() int {
	if t == nil {
		return NIL
	}

	return t.tType
}

func (t *T) GetVal() any {
	return t.val
}

func (t *T) GetDefineArgs() []string {
	return t.defineArgs
}

func (t *T) GetMethodName() string {
	return t.method
}

func (t *T) GetInfo() string {
	return t.GetFrame() + "::" + t.GetObjectClass() + "::" + t.method
}

func (t *T) GetObjectClass() string {
	return t.objectClass
}

func (t *T) SetObjectClass(class string) {
	t.objectClass = class
	t.val = class
}

func (t *T) GetKey() string {
	return t.key
}

func (t *T) GetFrame() string {
	return t.frame
}

func (t *T) SetFrame(frame string) {
	t.frame = frame
}

func (t *T) GetRemoveSuffixKey() string {
	return t.key[:len(t.key)-1]
}

func (t *T) GetRemovePrefixKey() string {
	return t.key[1:]
}

func (t *T) GetKeyValue() *T {
	return t.val.(*T)
}

// Block parameter accessors
func (t *T) SetBlockParamaters(blockParamaters []T) {
	t.blockParamaters = []T{}

	t.blockParamaters = append(t.blockParamaters, blockParamaters...)

	if len(t.blockParamaters) > 0 {
		t.IsBlockGiven = true
	}
}

func (t *T) GetBlockParameters() []T {
	if t == nil {
		return []T{}
	}

	return t.blockParamaters
}

// Flag setters
func (t *T) EnableReadOnly() {
	t.isReadOnly = true
}

func (t *T) DisableReadOnly() {
	t.isReadOnly = false
}

func (t *T) SetIsInfferedFromCall(b bool) {
	t.isInfferedFromCall = b
}

func (t *T) SetHasDefault(b bool) {
	t.hasDefault = b
}

// Code evaluation accessors
func (t *T) SetBeforeEvaluateCode(code string) {
	t.beforeEvaluateCode = code
}

func (t *T) GetBeforeEvaluateCode() string {
	return t.beforeEvaluateCode
}

// Variant and union type accessors
func (t *T) GetVariants() []T {
	return t.variants
}

func (t *T) PopVariants() *T {
	if len(t.variants) == 0 {
		return nil
	}

	returnT := t.variants[0].DeepCopy()

	t.variants = t.variants[1:]

	return returnT
}

func (t *T) GetVariantTypes() []int {
	var types []int
	for _, candidateT := range t.GetVariants() {
		types = append(types, candidateT.tType)
	}

	return types
}

// Variant manipulation functions
func (t *T) AppendVariant(variantT T) {
	switch variantT.tType {
	case UNION:
		for _, unionVariant := range variantT.variants {
			t.AppendVariant(unionVariant)
		}

		return

	case HASH:
		for _, variant := range t.variants {
			if variant.IsHashType() {
				variant.MergeHash(&variantT)
				return
			}
		}

		t.variants = append(t.variants, variantT)

		return

	case ARRAY:
		if len(t.variants) == 0 {
			t.variants = append(t.variants, variantT)
			return
		}

		var isArrayContained bool

		for currentTVariantIdx, currentTVariant := range t.variants {
			if !currentTVariant.IsArrayType() {
				continue
			}

			isArrayContained = true

			newVariants := currentTVariant.GetVariants()

			targetVariants := variantT.GetVariants()

			for targetTVariantIdx, targetTVariant := range targetVariants {
				if targetTVariantIdx >= len(newVariants) {
					newVariants = append(newVariants, targetTVariant)
				}

				if newVariants[targetTVariantIdx].IsEqualObject(&targetTVariant) {
					continue
				}

				if newVariants[targetTVariantIdx].IsUnionType() {
					newVariants[targetTVariantIdx].AppendVariant(targetTVariant)
					continue
				}

				unionT :=
					MakeUnifiedT([]T{newVariants[targetTVariantIdx], targetTVariant})

				newVariants[targetTVariantIdx] = *unionT
			}

			arrayT := MakeArray(newVariants)
			t.variants[currentTVariantIdx] = *arrayT
		}

		if !isArrayContained {
			t.variants = append(t.variants, variantT)
		}

	default:
		if !t.IsEqualObject(&variantT) {
			t.variants = append(t.variants, variantT)
		}
	}
}

func (t *T) AppendArrayVariant(appendT T) {
	t.variants = append(t.variants, appendT)
}

func (t *T) AppendHashVariant(keyvalueT T) {
	for idx, variant := range t.variants {
		if variant.key == keyvalueT.key {
			t.variants[idx] = keyvalueT
			return
		}
	}

	t.variants = append(t.variants, keyvalueT)
}

func (t *T) SetOwnerT(ownerT *T) {
	t.owner = ownerT
}

func (t *T) GetOwnerT() *T {
	if t.owner == nil {
		return nil
	}

	return t.owner.(*T)
}

func (t *T) ClearDefineArgs() {
	t.defineArgs = []string{}
}

func (t *T) EnableBuiltin() {
	t.isBuiltin = true
}
