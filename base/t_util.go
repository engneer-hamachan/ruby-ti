package base

import (
	"fmt"
	"strings"
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
	return t.ToString()[:len(t.ToString())-1]
}

func IsKeySuffix(str string) bool {
	return str[len(str)-1:] == ":" && len(str) >= 2
}

func RemoveSuffix(str string) string {
	return str[:len(str)-1]
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

	return MakeUntyped()
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

func SplitNameSpace(str string) []string {
	return strings.Split(str, "::")
}

func SeparateNameSpaces(str string) (string, string, string) {
	spaces := SplitNameSpace(str)

	// Hoge
	if len(spaces) == 1 {
		return "", "", spaces[0]
	}

	// Hoge::Fuga
	if len(spaces) == 2 {
		return "", spaces[0], spaces[1]
	}

	// Hoge::Fuga::Piyo
	frame := spaces[0]
	for _, name := range spaces[1 : len(spaces)-2] {
		frame += "::" + name
	}

	return frame, spaces[len(spaces)-2:][0], spaces[len(spaces)-2:][1]
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
