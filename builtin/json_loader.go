package builtin

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"path/filepath"
	"slices"
	"ti/base"
)

type MethodArgument struct {
	Type       []string       `json:"type,omitempty"`
	Key        string         `json:"key,omitempty"`
	IsAsterisk bool           `json:"is_asterisk,omitempty"`
	Value      map[string]any `json:",inline"`
}

type MethodReturn struct {
	Type                []string `json:"type,omitempty"`
	IsConditionalReturn bool     `json:"is_conditional,omitempty"`
	IsDestructive       bool     `json:"is_destructive,omitempty"`
}

type MethodDefinition struct {
	Name            string           `json:"name"`
	BlockParameters []string         `json:"block_parameters"`
	Arguments       []MethodArgument `json:"arguments"`
	ReturnType      MethodReturn     `json:"return_type"`
}

type ConstDefinition struct {
	Name       string       `json:"name"`
	ReturnType MethodReturn `json:"return_type"`
}

type ClassDefinition struct {
	Frame           string             `json:"frame"`
	Class           string             `json:"class"`
	InstanceMethods []MethodDefinition `json:"instance_methods"`
	ClassMethods    []MethodDefinition `json:"class_methods"`
	Constants       []ConstDefinition  `json:"constants"`
	Extends         []string           `json:"extends"`
}

func parseReturnType(returnType MethodReturn) base.T {
	var t base.T

	switch len(returnType.Type) {
	case 0:
		t = NilT

	case 1:
		t = parseTypeString(returnType.Type[0])

	default:
		var types []base.T

		for _, t := range returnType.Type {
			types = append(types, parseTypeString(t))
		}

		t = *base.MakeUnion(types)
	}

	t.IsConditionalReturn = returnType.IsConditionalReturn
	t.IsDestructive = returnType.IsDestructive

	return t
}

func parseTypeString(typeStr string) base.T {
	switch typeStr {
	case "Nil":
		return NilT
	case "Symbol":
		return SymbolT
	case "Bool":
		return BoolT
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
	case "SelfConvertArray":
		return SelfConvertArrayT
	case "SelfArgument":
		return SelfArgumentT
	case "KeyArray":
		return KeyArrayT
	case "KeyValueArray":
		return KeyValueArrayT
	case "UnifiedSelfArgument":
		return UnifiedSelfArgumentT
	default:
		return *base.MakeObject(typeStr)
	}
}

func parseArguments(args []MethodArgument) []base.T {
	var result []base.T

	for _, arg := range args {
		var baseType base.T

		switch len(arg.Type) {
		case 0:
			baseType = NilT

		case 1:
			baseType = parseTypeString(arg.Type[0])

		default:
			var types []base.T

			for _, t := range arg.Type {
				types = append(types, parseTypeString(t))
			}

			baseType = *base.MakeUnion(types)
		}

		baseType.IsBuiltinAsterisk = arg.IsAsterisk

		switch arg.Key {
		case "":
			result = append(result, baseType)

		default:
			keywordType := base.MakeKeyValue(arg.Key, &baseType)
			result = append(result, *keywordType)
		}
	}

	return result
}

func appendBlockParameters(returnType *base.T, method MethodDefinition) {
	var blockParameters []base.T

	for _, parameter := range method.BlockParameters {
		blockParameters = append(blockParameters, parseTypeString(parameter))
	}

	if len(blockParameters) > 0 {
		returnType.IsBlockGiven = true
	}

	returnType.SetBlockParamaters(blockParameters)
}

func loadBuiltinFromJSON(configFS fs.FS, configDir string) error {
	pattern := filepath.Join(configDir, "*.json")

	matches, err := fs.Glob(configFS, pattern)
	if err != nil {
		return fmt.Errorf("failed to find JSON config files: %w", err)
	}

	for _, match := range matches {
		jsonData, err := fs.ReadFile(configFS, match)
		if err != nil {
			return fmt.Errorf("failed to read %s: %w", match, err)
		}

		var classDef ClassDefinition
		if err := json.Unmarshal(jsonData, &classDef); err != nil {
			return fmt.Errorf("failed to parse %s: %w", match, err)
		}

		d := NewDefineBuiltinMethod(classDef.Frame, classDef.Class)

		base.BuiltinClasses = append(base.BuiltinClasses, classDef.Class)

		for _, method := range classDef.InstanceMethods {
			args := parseArguments(method.Arguments)
			returnType := parseReturnType(method.ReturnType)

			if len(method.BlockParameters) > 0 {
				appendBlockParameters(&returnType, method)
			}

			d.defineBuiltinInstanceMethod(
				classDef.Frame,
				method.Name,
				args,
				returnType,
			)
		}

		for _, method := range classDef.ClassMethods {
			args := parseArguments(method.Arguments)
			returnType := parseReturnType(method.ReturnType)

			if len(method.BlockParameters) > 0 {
				appendBlockParameters(&returnType, method)
			}

			d.defineBuiltinStaticMethod(
				classDef.Frame,
				method.Name,
				args,
				returnType,
			)
		}

		for _, constant := range classDef.Constants {
			returnType := parseReturnType(constant.ReturnType)

			d.defineBuiltinConstant(
				classDef.Frame,
				classDef.Class,
				constant.Name,
				returnType,
			)
		}

		// extends Object Class
		if classDef.Class != "" && classDef.Class != "Kernel" {
			classNode := base.ClassNode{Frame: classDef.Frame, Class: classDef.Class}
			parentNode := base.ClassNode{Frame: "Builtin", Class: ""}

			base.ClassInheritanceMap[classNode] =
				append(base.ClassInheritanceMap[classNode], parentNode)
		}

		// extends Other Class
		for _, class := range classDef.Extends {
			classNode := base.ClassNode{Frame: classDef.Frame, Class: classDef.Class}
			parentNode := base.ClassNode{Frame: classDef.Frame, Class: class}

			if slices.Contains(base.ClassInheritanceMap[classNode], parentNode) {
				continue
			}

			base.ClassInheritanceMap[classNode] =
				append(base.ClassInheritanceMap[classNode], parentNode)
		}
	}

	return nil
}
