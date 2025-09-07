package base

import "fmt"

const (
	EOS                   = -1
	NIL                   = 0
	INT                   = 257
	UNKNOWN               = 258
	STRING                = 259
	BOOL                  = 260
	FLOAT                 = 261
	UNTYPED               = 262
	ARRAY                 = 263
	HASH                  = 264
	UNION                 = 265
	OBJECT                = 266
	BLOCK                 = 267
	CLASS                 = 268
	SELF                  = 269
	SYMBOL                = 270
	KEYVALUE              = 271
	CONST                 = 272
	RANGE                 = 273
	UNIFY                 = 274
	OPTIONAL_UNIFY        = 275
	BLOCK_RESULT_ARRAY    = 276
	SELF_CONVERT_ARRAY    = 277
	SELF_ARGUMENT         = 278
	UNIFIED_SELF_ARGUMENT = 279
	KEYVALUE_ARRAY        = 280
)

func ArrayTypeToString(t *T) string {
	str := "Array<"

	variants := t.UnifyVariants().GetVariants()

	switch len(variants) {
	case 0:
		str += TypeToString(t.UnifyVariants())

	default:
		for _, variantT := range t.UnifyVariants().GetVariants() {
			switch variantT.GetType() {
			case UNION:
				str += UnionTypeToString(variantT.variants)

			default:
				str += TypeToString(&variantT)
			}

			str += " "
		}

		str = str[:len(str)-1]
	}

	str += ">"

	return str
}

func UnionTypeToString(unionTypes []T) string {
	str := "Union<"

	for _, variantT := range unionTypes {
		switch variantT.tType {
		case UNION:
			str += UnionTypeToString(variantT.GetVariants())

		default:
			str += TypeToString(&variantT)
		}

		str += " "
	}

	str = str[:len(str)-1]
	str += ">"

	return str
}

func TypeToString(t *T) string {
	if t == nil {
		return "?"
	}

	switch t.tType {
	case NIL:
		return "Nil"
	case INT:
		return "Integer"
	case UNKNOWN:
		return "Unknown"
	case STRING:
		return "String"
	case BOOL:
		return "Bool"
	case FLOAT:
		return "Float"
	case UNTYPED:
		return "untyped"
	case ARRAY:
		return ArrayTypeToString(t)
	case HASH:
		return "Hash"
	case UNION:
		return UnionTypeToString(t.variants)
	case OBJECT:
		return t.ToString()
	case BLOCK:
		return "Block"
	case CLASS:
		return t.ToString()
	case SELF:
		return "Self"
	case SYMBOL:
		return "Symbol"
	case KEYVALUE:
		return "Keyword"
	case CONST:
		return "Const"
	case RANGE:
		return "Range"
	case UNIFY:
		return "Unify"
	case OPTIONAL_UNIFY:
		return "OptiionalUnify"
	case SELF_CONVERT_ARRAY:
		return "SelfConvertArray"
	case SELF_ARGUMENT:
		return "SelfArgument"
	case UNIFIED_SELF_ARGUMENT:
		return "UnifiedSelfArgument"
	default:
		fmt.Println(t.tType)
		panic("type convert error")
	}
}
