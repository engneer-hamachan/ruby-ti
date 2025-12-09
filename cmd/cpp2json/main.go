package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type TiArgument struct {
	Type []string `json:"type"`
	Key  string   `json:"key,omitempty"`
}

type TiReturnType struct {
	Type []string `json:"type"`
}

type TiMethod struct {
	Name       string       `json:"name"`
	Arguments  []TiArgument `json:"arguments"`
	ReturnType TiReturnType `json:"return_type"`
	Document   string       `json:"document,omitempty"`
}

type TiConstantType struct {
	Name       string       `json:"name"`
	ReturnType TiReturnType `json:"return_type"`
}

type TiClassConfig struct {
	Frame           string           `json:"frame"`
	Class           string           `json:"class"`
	Extends         []string         `json:"extends,omitempty"`
	Constants       []TiConstantType `json:"constants,omitempty"`
	ClassMethods    []TiMethod       `json:"class_methods,omitempty"`
	InstanceMethods []TiMethod       `json:"instance_methods,omitempty"`
}

type FunctionInfo struct {
	Name string
	Body string
}

func isValidCmdArgumentCount() bool {
	flag.Parse()
	return flag.NArg() >= 1
}

func main() {
	output := flag.String("o", "", "output JSON file path")
	className := flag.String("class", "", "class or module name")
	isModule := flag.Bool("module", false, "define as module (use class_methods)")

	if !isValidCmdArgumentCount() {
		fmt.Fprintf(os.Stderr, "Usage: cpp2json [options] <input.cpp>\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	inputFile := flag.Arg(0)

	content, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	config := parseFile(string(content), *className, *isModule)

	jsonData, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating JSON: %v\n", err)
		os.Exit(1)
	}

	if *output != "" {
		if err := os.WriteFile(*output, jsonData, 0644); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing output file: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Generated: %s\n", *output)
	} else {
		fmt.Println(string(jsonData))
	}
}

func parseFile(content string, className string, isModule bool) TiClassConfig {
	config := TiClassConfig{
		Frame: "Builtin",
	}

	if className == "" {
		className = extractClassName(content)
	}
	config.Class = className

	functions := extractFunctions(content)
	methods := extractMethods(content)
	constants := extractConstants(content)

	config.Constants = constants

	for _, method := range methods {
		funcInfo, found := functions[method.FuncName]
		if !found {
			continue
		}

		methodDef := analyzeFunction(funcInfo, method.MethodName, method.ArgsSpec)

		if method.MethodName == "_init" {
			methodDef.Name = "new"
			config.ClassMethods = append(config.ClassMethods, methodDef)
		} else if isModule || method.MethodType == "class" {
			config.ClassMethods = append(config.ClassMethods, methodDef)
		} else {
			config.InstanceMethods = append(config.InstanceMethods, methodDef)
		}
	}

	return config
}

func extractClassName(content string) string {
	classOrModulePattern := regexp.MustCompile(`mrbc_define_(?:module|class)\s*\(\s*\w+\s*,\s*"([^"]+)"`)
	if matches := classOrModulePattern.FindStringSubmatch(content); matches != nil {
		return matches[1]
	}
	return "Unknown"
}

type MethodDefinition struct {
	MethodName string
	FuncName   string
	MethodType string
	ArgsSpec   string
}

func extractMethods(content string) []MethodDefinition {
	var methods []MethodDefinition

	mrbcDefinePattern := regexp.MustCompile(`mrbc_define_(method|class_method)\s*\(\s*\w+\s*,\s*\w+\s*,\s*"([^"]+)"\s*,\s*(\w+)\s*\)`)
	mrbcMatches := mrbcDefinePattern.FindAllStringSubmatch(content, -1)

	for _, matchGroups := range mrbcMatches {
		methodType := "instance"
		if matchGroups[1] == "class_method" {
			methodType = "class"
		}
		methods = append(methods, MethodDefinition{
			MethodName: matchGroups[2],
			FuncName:   matchGroups[3],
			MethodType: methodType,
		})
	}

	mrbDefineIdPattern := regexp.MustCompile(`mrb_define_(class_)?method_id\s*\(\s*\w+\s*,\s*\w+\s*,\s*MRB_SYM(_Q)?\((\w+)\)\s*,\s*(\w+)\s*,\s*([^)]+\))`)
	mrbIdMatches := mrbDefineIdPattern.FindAllStringSubmatch(content, -1)

	for _, matchGroups := range mrbIdMatches {
		methodType := "instance"
		if matchGroups[1] == "class_" {
			methodType = "class"
		}
		methodName := matchGroups[3]
		if matchGroups[2] == "_Q" {
			methodName += "?"
		}
		argumentsSpec := matchGroups[5]
		methods = append(methods, MethodDefinition{
			MethodName: methodName,
			FuncName:   matchGroups[4],
			MethodType: methodType,
			ArgsSpec:   argumentsSpec,
		})
	}

	mrbDefinePattern := regexp.MustCompile(`mrb_define_(class_)?method\s*\(\s*\w+\s*,\s*\w+\s*,\s*"([^"]+)"\s*,\s*(\w+)\s*,\s*([^)]+\))`)
	mrbMatches := mrbDefinePattern.FindAllStringSubmatch(content, -1)

	for _, matchGroups := range mrbMatches {
		methodType := "instance"
		if matchGroups[1] == "class_" {
			methodType = "class"
		}
		methodName := matchGroups[2]
		argumentsSpec := matchGroups[4]
		methods = append(methods, MethodDefinition{
			MethodName: methodName,
			FuncName:   matchGroups[3],
			MethodType: methodType,
			ArgsSpec:   argumentsSpec,
		})
	}

	return methods
}

func extractConstants(content string) []TiConstantType {
	var constants []TiConstantType

	constIdPattern := regexp.MustCompile(`mrb_define_const_id\s*\(\s*\w+\s*,\s*\w+\s*,\s*MRB_SYM\((\w+)\)\s*,\s*mrb_(\w+)_value\(`)
	constIdMatches := constIdPattern.FindAllStringSubmatch(content, -1)

	for _, matchGroups := range constIdMatches {
		constantName := matchGroups[1]
		mrubyValueType := matchGroups[2]

		rubyTiTypeName := convertMrubyValueTypeToRubyTiType(mrubyValueType)

		constants = append(constants, TiConstantType{
			Name: constantName,
			ReturnType: TiReturnType{
				Type: []string{rubyTiTypeName},
			},
		})
	}

	constPattern := regexp.MustCompile(`mrb_define_const\s*\(\s*\w+\s*,\s*\w+\s*,\s*"([^"]+)"\s*,\s*mrb_(\w+)_value\(`)
	constMatches := constPattern.FindAllStringSubmatch(content, -1)

	for _, matchGroups := range constMatches {
		constantName := matchGroups[1]
		mrubyValueType := matchGroups[2]

		rubyTiTypeName := convertMrubyValueTypeToRubyTiType(mrubyValueType)

		constants = append(constants, TiConstantType{
			Name: constantName,
			ReturnType: TiReturnType{
				Type: []string{rubyTiTypeName},
			},
		})
	}

	return constants
}

func convertMrubyValueTypeToRubyTiType(mrubyValueType string) string {
	switch mrubyValueType {
	case "fixnum", "int":
		return "Int"
	case "float":
		return "Float"
	case "true", "false":
		return "Bool"
	case "nil":
		return "Nil"
	case "str":
		return "String"
	case "symbol":
		return "Symbol"
	default:
		return "Untyped"
	}
}

func extractFunctions(content string) map[string]FunctionInfo {
	functionsByName := make(map[string]FunctionInfo)

	functionPattern := regexp.MustCompile(`(?s)(void|static\s+mrb_value)\s+(\w+)\s*\([^)]*\)\s*\{(.*?)\n\}`)
	allMatches := functionPattern.FindAllStringSubmatch(content, -1)

	for _, matchGroups := range allMatches {
		functionName := matchGroups[2]
		functionBody := matchGroups[3]
		functionsByName[functionName] = FunctionInfo{
			Name: functionName,
			Body: functionBody,
		}
	}

	return functionsByName
}

func analyzeFunction(functionInfo FunctionInfo, methodName string, argumentsSpec string) TiMethod {
	inferredReturnType := inferReturnType(functionInfo.Body, methodName)
	inferredArguments := inferArguments(functionInfo.Body, argumentsSpec)

	method := TiMethod{
		Name:       methodName,
		Arguments:  inferredArguments,
		ReturnType: TiReturnType{Type: []string{inferredReturnType}},
		Document:   "",
	}

	return method
}

func inferReturnType(functionBody string, methodName string) string {
	if strings.Contains(functionBody, "SET_NIL_RETURN()") ||
		strings.Contains(functionBody, "return mrb_nil_value()") ||
		strings.Contains(functionBody, "return mrbc_nil_value()") {
		return "Nil"
	}

	if strings.Contains(functionBody, "SET_INT_RETURN(") ||
		strings.Contains(functionBody, "return mrb_fixnum_value(") ||
		strings.Contains(functionBody, "return mrb_int_value(") ||
		strings.Contains(functionBody, "return mrbc_integer_value(") {
		return "Int"
	}

	if strings.Contains(functionBody, "SET_FLOAT_RETURN(") ||
		strings.Contains(functionBody, "return mrb_float_value(") ||
		strings.Contains(functionBody, "return mrbc_float_value(") {
		return "Float"
	}

	if strings.Contains(functionBody, "SET_TRUE_RETURN()") ||
		strings.Contains(functionBody, "SET_FALSE_RETURN()") ||
		strings.Contains(functionBody, "SET_BOOL_RETURN(") ||
		strings.Contains(functionBody, "return mrb_true_value()") ||
		strings.Contains(functionBody, "return mrb_false_value()") ||
		strings.Contains(functionBody, "return mrbc_true_value()") ||
		strings.Contains(functionBody, "return mrbc_false_value()") {
		return "Bool"
	}

	if strings.Contains(functionBody, "return mrb_str_new(") ||
		strings.Contains(functionBody, "return mrb_str_new_cstr(") ||
		strings.Contains(functionBody, "return mrb_str_new_lit(") ||
		strings.Contains(functionBody, "return mrbc_string_new(") ||
		strings.Contains(functionBody, "return mrbc_string_new_cstr(") {
		return "String"
	}

	if strings.Contains(functionBody, "return mrb_symbol_value(") ||
		strings.Contains(functionBody, "return mrbc_symbol_value(") {
		return "Symbol"
	}

	if strings.Contains(functionBody, "return mrb_ary_new(") ||
		strings.Contains(functionBody, "return mrb_ary_new_capa(") ||
		strings.Contains(functionBody, "return mrbc_array_new(") {
		return "Array"
	}

	if strings.Contains(functionBody, "return mrb_hash_new(") ||
		strings.Contains(functionBody, "return mrbc_hash_new(") {
		return "Hash"
	}

	instanceNewPattern := regexp.MustCompile(`mrbc_instance_new\s*\([^,]+,\s*mrbc_class_(\w+)`)
	if matches := instanceNewPattern.FindStringSubmatch(functionBody); matches != nil {
		rawClassName := matches[1]
		return normalizeClassName(rawClassName)
	}

	setReturnPattern := regexp.MustCompile(`SET_RETURN\s*\(\s*mrbc_instance_new\s*\([^,]+,\s*mrbc_class_(\w+)`)
	if matches := setReturnPattern.FindStringSubmatch(functionBody); matches != nil {
		rawClassName := matches[1]
		return normalizeClassName(rawClassName)
	}

	if strings.Contains(functionBody, "return self") ||
		strings.Contains(functionBody, "return mrb_obj_value(") {
		return "Self"
	}

	if strings.HasSuffix(methodName, "?") {
		return "Bool"
	}

	return "Untyped"
}

func normalizeClassName(rawClassName string) string {
	if len(rawClassName) == 0 {
		return rawClassName
	}
	if len(rawClassName) <= 2 {
		return strings.ToUpper(rawClassName)
	}
	return strings.ToUpper(rawClassName[:1]) + rawClassName[1:]
}

func inferArguments(functionBody string, argumentsSpec string) []TiArgument {
	arguments := []TiArgument{}

	if strings.Contains(argumentsSpec, "MRB_ARGS_NONE()") {
		return arguments
	}

	if strings.Contains(argumentsSpec, "MRB_ARGS_ANY()") {
		arguments = append(arguments, TiArgument{
			Type: []string{"Untyped"},
			Key:  "*args",
		})
		return arguments
	}

	requiredArgumentsCount := 0
	optionalArgumentsCount := 0
	hasRestArguments := false
	postArgumentsCount := 0
	hasBlockArgument := false

	requiredArgsPattern := regexp.MustCompile(`MRB_ARGS_REQ\((\d+)\)`)
	if matches := requiredArgsPattern.FindStringSubmatch(argumentsSpec); matches != nil {
		fmt.Sscanf(matches[1], "%d", &requiredArgumentsCount)
	}

	optionalArgsPattern := regexp.MustCompile(`MRB_ARGS_OPT\((\d+)\)`)
	if matches := optionalArgsPattern.FindStringSubmatch(argumentsSpec); matches != nil {
		fmt.Sscanf(matches[1], "%d", &optionalArgumentsCount)
	}

	if strings.Contains(argumentsSpec, "MRB_ARGS_REST()") {
		hasRestArguments = true
	}

	postArgsPattern := regexp.MustCompile(`MRB_ARGS_POST\((\d+)\)`)
	if matches := postArgsPattern.FindStringSubmatch(argumentsSpec); matches != nil {
		fmt.Sscanf(matches[1], "%d", &postArgumentsCount)
	}

	if strings.Contains(argumentsSpec, "MRB_ARGS_BLOCK()") {
		hasBlockArgument = true
	}

	getArgsPattern := regexp.MustCompile(`mrb_get_args\s*\(\s*\w+\s*,\s*"([^"]+)"`)
	if matches := getArgsPattern.FindStringSubmatch(functionBody); matches != nil {
		formatString := matches[1]
		isInOptionalSection := false

		for _, formatCharacter := range formatString {
			switch formatCharacter {
			case 'i':
				argumentType := "Int"
				if isInOptionalSection {
					argumentType = "DefaultInt"
				}
				arguments = append(arguments, TiArgument{Type: []string{argumentType}})
			case 'f':
				argumentType := "Float"
				if isInOptionalSection {
					argumentType = "DefaultFloat"
				}
				arguments = append(arguments, TiArgument{Type: []string{argumentType}})
			case 's', 'z':
				argumentType := "String"
				if isInOptionalSection {
					argumentType = "DefaultString"
				}
				arguments = append(arguments, TiArgument{Type: []string{argumentType}})
			case 'S':
				argumentType := "String"
				if isInOptionalSection {
					argumentType = "DefaultString"
				}
				arguments = append(arguments, TiArgument{Type: []string{argumentType}})
			case 'A', 'a':
				argumentType := "Array"
				if isInOptionalSection {
					argumentType = "DefaultUntyped"
				}
				arguments = append(arguments, TiArgument{Type: []string{argumentType}})
			case 'H':
				argumentType := "Hash"
				if isInOptionalSection {
					argumentType = "DefaultUntyped"
				}
				arguments = append(arguments, TiArgument{Type: []string{argumentType}})
			case 'b':
				argumentType := "Bool"
				if isInOptionalSection {
					argumentType = "DefaultUntyped"
				}
				arguments = append(arguments, TiArgument{Type: []string{argumentType}})
			case 'n':
				argumentType := "Symbol"
				if isInOptionalSection {
					argumentType = "DefaultUntyped"
				}
				arguments = append(arguments, TiArgument{Type: []string{argumentType}})
			case 'C', 'o':
				argumentType := "Untyped"
				if isInOptionalSection {
					argumentType = "DefaultUntyped"
				}
				arguments = append(arguments, TiArgument{Type: []string{argumentType}})
			case '|':
				isInOptionalSection = true
			case '*':
				arguments = append(arguments, TiArgument{Type: []string{"Untyped"}, Key: "*args"})
			case '&':
				arguments = append(arguments, TiArgument{Type: []string{"DefaultBlock"}})
			case '!', '?':
			}
		}
		return arguments
	}

	hasNoArgumentSpec := requiredArgumentsCount == 0 && optionalArgumentsCount == 0 && !hasRestArguments && postArgumentsCount == 0 && !hasBlockArgument
	if hasNoArgumentSpec {
		getArgTypePattern := regexp.MustCompile(`GET_(\w+)_ARG\s*\(\s*(\d+)\s*\)`)
		getArgTypeMatches := getArgTypePattern.FindAllStringSubmatch(functionBody, -1)
		argumentTypesByIndex := make(map[int]string)

		for _, matchGroups := range getArgTypeMatches {
			mrubyArgumentType := matchGroups[1]
			argumentIndex := 0
			fmt.Sscanf(matchGroups[2], "%d", &argumentIndex)

			switch mrubyArgumentType {
			case "INT":
				argumentTypesByIndex[argumentIndex] = "Int"
			case "FLOAT":
				argumentTypesByIndex[argumentIndex] = "Float"
			case "STRING":
				argumentTypesByIndex[argumentIndex] = "String"
			default:
				argumentTypesByIndex[argumentIndex] = "Untyped"
			}
		}

		argcCheckPattern := regexp.MustCompile(`if\s*\(\s*argc\s*>=\s*(\d+)\s*\)`)
		argcCheckMatches := argcCheckPattern.FindAllStringSubmatch(functionBody, -1)
		minimumRequiredArgc := 0
		for _, matchGroups := range argcCheckMatches {
			argcValue := 0
			fmt.Sscanf(matchGroups[1], "%d", &argcValue)
			if argcValue > minimumRequiredArgc {
				minimumRequiredArgc = argcValue
			}
		}

		if len(argumentTypesByIndex) > 0 {
			maxArgumentIndex := 0
			for argumentIndex := range argumentTypesByIndex {
				if argumentIndex > maxArgumentIndex {
					maxArgumentIndex = argumentIndex
				}
			}

			for currentIndex := 1; currentIndex <= maxArgumentIndex; currentIndex++ {
				argumentType := "Untyped"
				if inferredType, exists := argumentTypesByIndex[currentIndex]; exists {
					argumentType = inferredType
				}

				if minimumRequiredArgc > 0 && currentIndex >= minimumRequiredArgc {
					if argumentType == "Untyped" {
						argumentType = "DefaultUntyped"
					} else {
						argumentType = "Default" + argumentType
					}
				}

				arguments = append(arguments, TiArgument{
					Type: []string{argumentType},
				})
			}
			return arguments
		}
	}

	for i := 0; i < requiredArgumentsCount; i++ {
		arguments = append(arguments, TiArgument{
			Type: []string{"Untyped"},
		})
	}

	for i := 0; i < optionalArgumentsCount; i++ {
		arguments = append(arguments, TiArgument{
			Type: []string{"DefaultUntyped"},
		})
	}

	if hasRestArguments {
		arguments = append(arguments, TiArgument{
			Type: []string{"Untyped"},
			Key:  "*args",
		})
	}

	for i := 0; i < postArgumentsCount; i++ {
		arguments = append(arguments, TiArgument{
			Type: []string{"Untyped"},
		})
	}

	if hasBlockArgument {
		arguments = append(arguments, TiArgument{
			Type: []string{"DefaultBlock"},
		})
	}

	return arguments
}
