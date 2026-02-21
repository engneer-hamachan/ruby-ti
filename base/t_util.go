package base

import (
	"fmt"
)

// String utility functions
func (t *T) ToString() string {
	if t == nil {
		return ""
	}

	switch t.val.(type) {
	case string:
		return t.val.(string)
	case int64:
		return "Integer"
	case float64:
		return "Float"
	default:
		return "Unknown"
	}
}

func (t *T) ToRemoveSuffixString() string {
	return RemoveSuffix(t.ToString())
}

// Type conversion functions
func (t *T) UnifyVariants() *T {
	unionT := MakeUnion([]T{})

	switch t.tType {
	case HASH:
		for _, variantT := range t.variants {
			unionT.AppendVariant(*variantT.GetKeyValue())
		}

	default:
		for _, variantT := range t.variants {
			unionT.AppendVariant(variantT)
		}
	}

	if len(unionT.variants) == 0 {
		return MakeUntyped()
	}

	if len(unionT.variants) == 1 {
		return &unionT.variants[0]
	}

	return unionT
}

// Hash utility functions
func (t *T) HashReference(key string) *T {
	for _, variant := range t.variants {
		if variant.key == key {
			return variant.GetKeyValue()
		}
	}

	return t.UnifyVariants()
}

func (t *T) MergeHash(variantT *T) {
	if t.tType != HASH || variantT.tType != HASH {
		return
	}

	for _, variant := range variantT.variants {
		existT := t.HashReference(variant.key)
		newValueT := variant.GetKeyValue()

		switch existT.tType {
		case NIL:
			t.AppendHashVariant(variant)
			continue

		case UNION:
			if !existT.IsMatchUnionType(newValueT) {
				existT.AppendVariant(*newValueT)
			}

			continue

		case newValueT.tType:
			continue

		default:
			unionType := MakeUnion([]T{*existT, *newValueT})
			mergedVariant := MakeKeyValue(variant.key, unionType)
			t.AppendHashVariant(*mergedVariant)
		}
	}
}

type genId struct {
	prefix string
	count  int
}

var defaultGenId = genId{prefix: "var", count: 0}

func GenId() string {
	id := fmt.Sprintf("%s%d", defaultGenId.prefix, defaultGenId.count)

	defaultGenId.count++

	return id
}
