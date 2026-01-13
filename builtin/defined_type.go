package builtin

import (
	"strings"
	"ti/base"
)

var (
	//Other
	NilT              = *base.MakeNil()
	SymbolT           = *base.MakeAnySymbol()
	BoolT             = *base.MakeBool()
	DefaultBoolT      = *base.MakeBuiltinDefaultBool()
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
	SelfT          = *base.MakeSelf()
	NumberT        = *base.MakeUnion([]base.T{IntT, FloatT})
	UnifyT         = *base.MakeUnify()
	OptionalUnifyT = *base.MakeOptionalUnify()
	SelfArrayT     = *base.MakeSelfArray()
	ArgumentT      = *base.MakeArgument()
	UnifyArgumentT = *base.MakeUnifyArgument()
	FlattenT       = *base.MakeFlatten()
	ItemT          = *base.MakeItem()
	OwnerT         = *base.MakeOwner()
	//ForTest
	IntIntT = *base.MakeUnion([]base.T{IntT, IntT})
)

func ConvertToBuiltinT(typeStr string) base.T {
	switch typeStr {
	case "NilClass":
		return NilT
	case "Symbol":
		return SymbolT
	case "Bool":
		return BoolT
	case "DefaultBool":
		return DefaultBoolT
	case "Block":
		return BlockT
	case "DefaultBlock":
		return DefaultBlockT
	case "Range":
		return RangeT
	case "Untyped":
		return UntypedT
	case "DefaultUntyped":
		return DefaultUntypedT
	case "String":
		return StringT
	case "DefaultString":
		return DefaultStringT
	case "OptionalString":
		return OptionalStringT
	case "Int":
		return IntT
	case "DefaultInt":
		return DefaultIntT
	case "OptionalInt":
		return OptionalIntT
	case "Float":
		return FloatT
	case "DefaultFloat":
		return DefaultFloatT
	case "OptionalFloat":
		return OptionalFloatT
	case "Array":
		return ArrayT
	case "Hash":
		return HashT
	case "StringArray":
		return StringArrayT
	case "IntArray":
		return IntArrayT
	case "FloatArray":
		return FloatArrayT
	case "Self":
		return SelfT
	case "Number":
		return NumberT
	case "IntInt":
		return IntIntT
	case "Unify":
		return UnifyT
	case "OptionalUnify":
		return OptionalUnifyT
	case "BlockResultArray":
		return BlockResultArrayT
	case "SelfArray":
		return SelfArrayT
	case "Argument":
		return ArgumentT
	case "KeyArray":
		return KeyArrayT
	case "KeyValueArray":
		return KeyValueArrayT
	case "UnifyArgument":
		return UnifyArgumentT
	case "Flatten":
		return FlattenT
	case "Item":
		return ItemT
	case "Owner":
		return OwnerT
	default:
		if len(strings.Split(typeStr, "::")) > 1 {
			return *base.MakeIdentifier(typeStr)
		}

		return *base.MakeObject(typeStr)
	}
}
