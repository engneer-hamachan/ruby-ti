package base

func NewT(objectClass string, types int, val any) *T {
	return &T{
		objectClass: objectClass,
		tType:       types,
		val:         val,
	}
}

func MakeObjectObject() *T {
	t := NewT("", OBJECT, "")

	return t
}

func MakeInt(val int64) *T {
	t := NewT("Integer", INT, val)

	return t
}

func MakeAnyInt() *T {
	t := NewT("Integer", INT, 1)

	return t
}

func MakeBuiltinDefaultInt() *T {
	t := NewT("Integer", INT, 1)
	t.hasDefault = true
	t.isBuiltin = true

	return t
}

func MakeFloat(val float64) *T {
	return NewT("Float", FLOAT, val)
}

func MakeAnyFloat() *T {
	t := NewT("Float", FLOAT, 1)

	return t
}

func MakeBuiltinDefaultFloat() *T {
	t := NewT("Float", FLOAT, 1)
	t.hasDefault = true
	t.isBuiltin = true

	return t
}

func MakeString(val string) *T {
	return NewT("String", STRING, val)
}

func MakeAnyString() *T {
	t := NewT("String", STRING, "String")

	return t
}

func MakeBuiltinDefaultString() *T {
	t := NewT("String", STRING, "String")
	t.hasDefault = true
	t.isBuiltin = true

	return t
}

func MakeArray() *T {
	return NewT("Array", ARRAY, "array")
}

func MakeStringArray() *T {
	t := NewT("Array", ARRAY, "array")
	t.AppendArrayVariant(*MakeAnyString())

	return t
}

func MakeIntArray() *T {
	t := NewT("Array", ARRAY, "array")
	t.AppendArrayVariant(*MakeAnyInt())

	return t
}

func MakeFloatArray() *T {
	t := NewT("Array", ARRAY, "array")
	t.AppendArrayVariant(*MakeAnyFloat())

	return t
}

func MakeAnyArray() *T {
	t := NewT("Array", ARRAY, "array")

	return t
}

func MakeHash(val string) *T {
	return NewT("Hash", HASH, val)
}

func MakeAnyHash() *T {
	t := NewT("Hash", HASH, "hash")

	return t
}

func MakeRange() *T {
	return NewT("Range", RANGE, "range")
}

func MakeBool() *T {
	return NewT("Bool", BOOL, "bool")
}

func MakeNil() *T {
	return NewT("Nil", NIL, "nil")
}

func MakeIdentifier(val string) *T {
	id := NewT("Identifier", UNKNOWN, val)

	return id
}

func MakeUnknown() *T {
	return NewT("", UNKNOWN, "")
}

func MakeObject(val string) *T {
	return NewT(val, OBJECT, val)
}

func MakeClass(val string) *T {
	return NewT(val, CLASS, val)
}

func MakeConst(val string) *T {
	return NewT(val, CONST, val)
}

func MakeUnion(variants []T) *T {
	t := NewT("Union", UNION, "union")

	t.variants = append(t.variants, variants...)

	return t
}

func MakeUnifiedT(variants []T) *T {
	t := NewT("Union", UNION, "union")
	t.variants = append(t.variants, variants...)

	return t.UnifyVariants()
}

func MakeBlockWithResult(val *T) *T {
	t := NewT("Block", BLOCK, "block")
	t.val = val

	return t
}

func MakeBlock() *T {
	return NewT("Block", BLOCK, "block")
}

func MakeBuiltinDefaultBlock() *T {
	t := NewT("Block", BLOCK, "block")
	t.hasDefault = true
	t.isBuiltin = true

	return t
}

func MakeUntyped() *T {
	return NewT("Untyped", UNTYPED, "untyped")
}

func MakeBuiltinDefaultUntyped() *T {
	t := NewT("Untyped", UNTYPED, "untyped")
	t.hasDefault = true
	t.isBuiltin = true

	return t
}

func MakeSelf() *T {
	return NewT("Self", SELF, "self")
}

func MakeUnify() *T {
	return NewT("Unify", UNIFY, "unify")
}

func MakeOptionalUnify() *T {
	return NewT("OptiionalUnify", OPTIONAL_UNIFY, "optionalUnify")
}

func MakeSelfConvertArray() *T {
	return NewT("SelfConvertArray", SELF_CONVERT_ARRAY, "selfConvertArray")
}

func MakeSelfArgument() *T {
	return NewT("SelfArgument", SELF_ARGUMENT, "selfArgument")
}

func MakeUnifiedSelfArgument() *T {
	return NewT(
		"UnifiedSelfArgument",
		UNIFIED_SELF_ARGUMENT,
		"unifiedSelfArgument",
	)
}

func MakeSymbol(val string) *T {
	return NewT("Symbol", SYMBOL, val)
}

func MakeAnySymbol() *T {
	return NewT("Symbol", SYMBOL, "symbol")
}

func MakeKeyValue(key string, valueT *T) *T {
	t := NewT("KeyValue", KEYVALUE, valueT)
	t.key = key

	return t
}

func MakeMethod(
	frame string,
	method string,
	returnT T,
	args []string,
) *T {

	returnT.frame = frame
	returnT.method = method
	returnT.defineArgs = args

	return &returnT
}

func MakeBlockResultArray() *T {
	return NewT("BlockResultArray", BLOCK_RESULT_ARRAY, "blockResultArray")
}

func MakeKeyArray() *T {
	t := NewT("Array", ARRAY, "array")
	unionT := MakeUnion([]T{*MakeAnyString(), *MakeAnySymbol()})

	t.AppendArrayVariant(*unionT)

	return t
}

func MakeKeyValueArray() *T {
	t := NewT("KeyValueArray", KEYVALUE_ARRAY, "keyValueArray")

	return t
}

func MakeEos() *T {
	return NewT("EOS", EOS, "eos")
}
