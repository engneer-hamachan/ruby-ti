package builtin

import "ti/base"

var (
	//Other
	NilT              = *base.MakeNil()
	SymbolT           = *base.MakeAnySymbol()
	BoolT             = *base.MakeBool()
	RangeT            = *base.MakeRange()
	BlockResultArrayT = *base.MakeBlockResultArray()
	//Block
	BlockT        = *base.MakeBlock()
	DefaultBlockT = *base.MakeBuiltinDefaultBlock()
	//Any
	UntypedT        = *base.MakeUntyped()
	DefaultUntypedT = *base.MakeBuiltinDefaultUntyped()
	//String
	StringT         = *base.MakeAnyString()
	DefaultStringT  = *base.MakeBuiltinDefaultString()
	OptionalStringT = *base.MakeUnion([]base.T{StringT, NilT})
	//Int
	IntT         = *base.MakeAnyInt()
	DefaultIntT  = *base.MakeBuiltinDefaultInt()
	OptionalIntT = *base.MakeUnion([]base.T{IntT, NilT})
	//Float
	FloatT         = *base.MakeAnyFloat()
	DefaultFloatT  = *base.MakeBuiltinDefaultFloat()
	OptionalFloatT = *base.MakeUnion([]base.T{FloatT, NilT})
	//Array
	ArrayT       = *base.MakeAnyArray()
	StringArrayT = *base.MakeStringArray()
	IntArrayT    = *base.MakeIntArray()
	FloatArrayT  = *base.MakeFloatArray()
	//Hash
	HashT          = *base.MakeAnyHash()
	KeyArrayT      = *base.MakeKeyArray()
	KeyValueArrayT = *base.MakeKeyValueArray()
	//Special
	SelfT                = *base.MakeSelf()
	NumberT              = *base.MakeUnion([]base.T{IntT, FloatT})
	UnifyT               = *base.MakeUnify()
	OptionalUnifyT       = *base.MakeOptionalUnify()
	SelfConvertArrayT    = *base.MakeSelfConvertArray()
	SelfArgumentT        = *base.MakeSelfArgument()
	UnifiedSelfArgumentT = *base.MakeUnifiedSelfArgument()
	FlattenT             = *base.MakeFlatten()
	ItemT                = *base.MakeItem()
	OwnerT               = *base.MakeOwner()
	//ForTest
	IntIntT = *base.MakeUnion([]base.T{IntT, IntT})
)
