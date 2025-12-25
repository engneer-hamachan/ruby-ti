package builtin

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"ti/base"
)

type MethodArgument struct {
	Type       TypeSpec       `json:"type,omitempty"`
	Key        string         `json:"key,omitempty"`
	IsAsterisk bool           `json:"is_asterisk,omitempty"`
	Value      map[string]any `json:",inline"`
}

type MethodReturn struct {
	Type                TypeSpec `json:"type,omitempty"`
	IsConditionalReturn bool     `json:"is_conditional,omitempty"`
	IsDestructive       bool     `json:"is_destructive,omitempty"`
	IsCaptureOwner      bool     `json:"is_capture_owner,omitempty"`
}

// TypeSpec handles both string and array type specifications
type TypeSpec []string

func (ts *TypeSpec) UnmarshalJSON(data []byte) error {
	// Try to unmarshal as a string first
	var single string
	if err := json.Unmarshal(data, &single); err == nil {
		*ts = []string{single}
		return nil
	}

	// If that fails, try as an array
	var array []string
	if err := json.Unmarshal(data, &array); err != nil {
		return err
	}
	*ts = array
	return nil
}

type MethodDefinition struct {
	Name            string           `json:"name"`
	BlockParameters []string         `json:"block_parameters"`
	Arguments       []MethodArgument `json:"arguments"`
	ReturnType      MethodReturn     `json:"return_type"`
	Document        string           `json:"document"`
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
	t.IsCaptureOwner = returnType.IsCaptureOwner

	return t
}

var AllTypeNames = []string{
	"Nil",
	"Symbol",
	"Bool",
	"DefaultBool",
	"Block",
	"DefaultBlock",
	"Range",
	"Untyped",
	"DefaultUntyped",
	"String",
	"DefaultString",
	"OptionalString",
	"Int",
	"DefaultInt",
	"OptionalInt",
	"Float",
	"DefaultFloat",
	"OptionalFloat",
	"Array",
	"Hash",
	"StringArray",
	"IntArray",
	"FloatArray",
	"Self",
	"Number",
	"IntInt",
	"Unify",
	"OptionalUnify",
	"BlockResultArray",
	"SelfArray",
	"Argument",
	"KeyArray",
	"KeyValueArray",
	"UnifyArgument",
	"Flatten",
	"Item",
	"Owner",
	"SymbolToMethod",
	"SymbolToMethods",
}

func parseTypeString(typeStr string) base.T {
	// Handle compact notation: ?Type, *Type, [Type], Type|Other

	// Optional type: ?String -> Union<String, Nil>
	if len(typeStr) > 1 && typeStr[0] == '?' {
		innerType := parseTypeString(typeStr[1:])
		return *base.MakeUnion([]base.T{innerType, NilT})
	}

	// Asterisk type: *String -> String with is_asterisk
	if len(typeStr) > 1 && typeStr[0] == '*' {
		innerType := parseTypeString(typeStr[1:])
		innerType.IsBuiltinAsterisk = true
		return innerType
	}

	// Array type: [String] -> Array<String>
	if len(typeStr) > 2 && typeStr[0] == '[' && typeStr[len(typeStr)-1] == ']' {
		innerType := typeStr[1 : len(typeStr)-1]
		innerT := parseTypeString(innerType)

		arrayT := *base.MakeArray([]base.T{innerT})

		return arrayT
	}

	// Union type: String|Int -> Union<String, Int>
	if strings.Contains(typeStr, "|") {
		parts := strings.Split(typeStr, "|")
		var types []base.T
		for _, part := range parts {
			types = append(types, parseTypeString(strings.TrimSpace(part)))
		}
		return *base.MakeUnion(types)
	}

	// Standard type names
	switch typeStr {
	case "Nil":
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

func parseArguments(args []MethodArgument) []base.T {
	var result []base.T

	for _, arg := range args {
		var baseType base.T

		switch len(arg.Type) {
		case 0:
			baseType = NilT

		case 1:
			typeStr := arg.Type[0]

			// Check for old-style prefix notation (for backward compatibility)
			if len(typeStr) > 0 {
				switch typeStr[0] {
				case '*':
					// Old style: "*String" in JSON means asterisk
					if !strings.Contains(typeStr, "|") && !strings.Contains(typeStr, "[") {
						arg.IsAsterisk = true
						typeStr = typeStr[1:]
						baseType = parseTypeString(typeStr)
					} else {
						baseType = parseTypeString(typeStr)
					}

				case '?':
					// Old style: "?String" in arguments means default parameter
					if !strings.Contains(typeStr, "|") && !strings.Contains(typeStr, "[") {
						typeStr = typeStr[1:]
						baseType = parseTypeString(typeStr)
						baseType.SetHasDefault(true)
					} else {
						baseType = parseTypeString(typeStr)
					}

				default:
					baseType = parseTypeString(typeStr)
				}
			} else {
				baseType = parseTypeString(typeStr)
			}

		default:
			// Multiple types specified as array: ["String", "Int"]
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

func init() {
	err := loadBuiltinFromJSON()
	if err != nil {
		panic("json loading error!")
	}
}

func loadBuiltinFromJSON() error {
	configDir := ".ti-config"
	pattern := filepath.Join(configDir, "*.json")

	matches, err := filepath.Glob(pattern)
	if err != nil {
		return fmt.Errorf("failed to find JSON config files in %s: %w", configDir, err)
	}

	for _, match := range matches {
		jsonData, err := os.ReadFile(match)
		if err != nil {
			return fmt.Errorf("failed to read %s: %w", match, err)
		}

		if len(jsonData) == 0 {
			continue
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

			key := classDef.Frame + classDef.Class + method.Name
			base.TSignatureDocument[key] = method.Document

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

			key := classDef.Frame + classDef.Class + method.Name + "static"
			base.TSignatureDocument[key] = method.Document

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

			var parentNode base.ClassNode

			parts := strings.Split(class, "::")

			switch len(parts) {
			case 2:
				parentNode = base.ClassNode{Frame: parts[0], Class: parts[1]}
			default:
				parentNode = base.ClassNode{Frame: classDef.Frame, Class: class}
			}

			if slices.Contains(base.ClassInheritanceMap[classNode], parentNode) {
				continue
			}

			base.ClassInheritanceMap[classNode] =
				append(base.ClassInheritanceMap[classNode], parentNode)
		}

		d.SetDefinedClass()
	}

	return nil
}
