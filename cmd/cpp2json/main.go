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

func main() {
	output := flag.String("o", "", "output JSON file path")
	className := flag.String("class", "", "class or module name")
	isModule := flag.Bool("module", false, "define as module (use class_methods)")
	flag.Parse()

	if flag.NArg() < 1 {
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
	moduleRe := regexp.MustCompile(`mrbc_define_(?:module|class)\s*\(\s*\w+\s*,\s*"([^"]+)"`)
	if match := moduleRe.FindStringSubmatch(content); match != nil {
		return match[1]
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

	mrbcDefineRe := regexp.MustCompile(`mrbc_define_(method|class_method)\s*\(\s*\w+\s*,\s*\w+\s*,\s*"([^"]+)"\s*,\s*(\w+)\s*\)`)
	matches := mrbcDefineRe.FindAllStringSubmatch(content, -1)

	for _, match := range matches {
		methodType := "instance"
		if match[1] == "class_method" {
			methodType = "class"
		}
		methods = append(methods, MethodDefinition{
			MethodName: match[2],
			FuncName:   match[3],
			MethodType: methodType,
		})
	}

	mrbDefineIdRe := regexp.MustCompile(`mrb_define_(class_)?method_id\s*\(\s*\w+\s*,\s*\w+\s*,\s*MRB_SYM(_Q)?\((\w+)\)\s*,\s*(\w+)\s*,\s*([^)]+\))`)
	mrbIdMatches := mrbDefineIdRe.FindAllStringSubmatch(content, -1)

	for _, match := range mrbIdMatches {
		methodType := "instance"
		if match[1] == "class_" {
			methodType = "class"
		}
		methodName := match[3]
		if match[2] == "_Q" {
			methodName += "?"
		}
		argsSpec := match[5]
		methods = append(methods, MethodDefinition{
			MethodName: methodName,
			FuncName:   match[4],
			MethodType: methodType,
			ArgsSpec:   argsSpec,
		})
	}

	mrbDefineRe := regexp.MustCompile(`mrb_define_(class_)?method\s*\(\s*\w+\s*,\s*\w+\s*,\s*"([^"]+)"\s*,\s*(\w+)\s*,\s*([^)]+\))`)
	mrbMatches := mrbDefineRe.FindAllStringSubmatch(content, -1)

	for _, match := range mrbMatches {
		methodType := "instance"
		if match[1] == "class_" {
			methodType = "class"
		}
		methodName := match[2]
		argsSpec := match[4]
		methods = append(methods, MethodDefinition{
			MethodName: methodName,
			FuncName:   match[3],
			MethodType: methodType,
			ArgsSpec:   argsSpec,
		})
	}

	return methods
}

func extractConstants(content string) []TiConstantType {
	var constants []TiConstantType

	constIdRe := regexp.MustCompile(`mrb_define_const_id\s*\(\s*\w+\s*,\s*\w+\s*,\s*MRB_SYM\((\w+)\)\s*,\s*mrb_(\w+)_value\(`)
	constIdMatches := constIdRe.FindAllStringSubmatch(content, -1)

	for _, match := range constIdMatches {
		constName := match[1]
		valueType := match[2]

		typeName := "Untyped"
		switch valueType {
		case "fixnum", "int":
			typeName = "Int"
		case "float":
			typeName = "Float"
		case "true", "false":
			typeName = "Bool"
		case "nil":
			typeName = "Nil"
		case "str":
			typeName = "String"
		case "symbol":
			typeName = "Symbol"
		}

		constants = append(constants, TiConstantType{
			Name: constName,
			ReturnType: TiReturnType{
				Type: []string{typeName},
			},
		})
	}

	constRe := regexp.MustCompile(`mrb_define_const\s*\(\s*\w+\s*,\s*\w+\s*,\s*"([^"]+)"\s*,\s*mrb_(\w+)_value\(`)
	constMatches := constRe.FindAllStringSubmatch(content, -1)

	for _, match := range constMatches {
		constName := match[1]
		valueType := match[2]

		typeName := "Untyped"
		switch valueType {
		case "fixnum", "int":
			typeName = "Int"
		case "float":
			typeName = "Float"
		case "true", "false":
			typeName = "Bool"
		case "nil":
			typeName = "Nil"
		case "str":
			typeName = "String"
		case "symbol":
			typeName = "Symbol"
		}

		constants = append(constants, TiConstantType{
			Name: constName,
			ReturnType: TiReturnType{
				Type: []string{typeName},
			},
		})
	}

	return constants
}

func extractFunctions(content string) map[string]FunctionInfo {
	functions := make(map[string]FunctionInfo)

	funcRe := regexp.MustCompile(`(?s)(void|static\s+mrb_value)\s+(\w+)\s*\([^)]*\)\s*\{(.*?)\n\}`)
	matches := funcRe.FindAllStringSubmatch(content, -1)

	for _, match := range matches {
		funcName := match[2]
		funcBody := match[3]
		functions[funcName] = FunctionInfo{
			Name: funcName,
			Body: funcBody,
		}
	}

	return functions
}

func analyzeFunction(funcInfo FunctionInfo, methodName string, argsSpec string) TiMethod {
	method := TiMethod{
		Name:      methodName,
		Arguments: []TiArgument{},
	}

	returnType := inferReturnType(funcInfo.Body, methodName)
	method.ReturnType = TiReturnType{Type: []string{returnType}}

	arguments := inferArguments(funcInfo.Body, argsSpec)
	method.Arguments = arguments
	method.Document = ""

	return method
}

func inferReturnType(body string, methodName string) string {
	if strings.Contains(body, "SET_NIL_RETURN()") || strings.Contains(body, "return mrb_nil_value()") ||
	   strings.Contains(body, "return mrbc_nil_value()") {
		return "Nil"
	}
	if strings.Contains(body, "SET_INT_RETURN(") || strings.Contains(body, "return mrb_fixnum_value(") ||
	   strings.Contains(body, "return mrb_int_value(") || strings.Contains(body, "return mrbc_integer_value(") {
		return "Int"
	}
	if strings.Contains(body, "SET_FLOAT_RETURN(") || strings.Contains(body, "return mrb_float_value(") ||
	   strings.Contains(body, "return mrbc_float_value(") {
		return "Float"
	}
	if strings.Contains(body, "SET_TRUE_RETURN()") || strings.Contains(body, "SET_FALSE_RETURN()") ||
	   strings.Contains(body, "SET_BOOL_RETURN(") || strings.Contains(body, "return mrb_true_value()") ||
	   strings.Contains(body, "return mrb_false_value()") || strings.Contains(body, "return mrbc_true_value()") ||
	   strings.Contains(body, "return mrbc_false_value()") {
		return "Bool"
	}
	if strings.Contains(body, "return mrb_str_new(") || strings.Contains(body, "return mrb_str_new_cstr(") ||
	   strings.Contains(body, "return mrb_str_new_lit(") || strings.Contains(body, "return mrbc_string_new(") ||
	   strings.Contains(body, "return mrbc_string_new_cstr(") {
		return "String"
	}
	if strings.Contains(body, "return mrb_symbol_value(") || strings.Contains(body, "return mrbc_symbol_value(") {
		return "Symbol"
	}
	if strings.Contains(body, "return mrb_ary_new(") || strings.Contains(body, "return mrb_ary_new_capa(") ||
	   strings.Contains(body, "return mrbc_array_new(") {
		return "Array"
	}
	if strings.Contains(body, "return mrb_hash_new(") || strings.Contains(body, "return mrbc_hash_new(") {
		return "Hash"
	}

	instanceNewRe := regexp.MustCompile(`mrbc_instance_new\s*\([^,]+,\s*mrbc_class_(\w+)`)
	if match := instanceNewRe.FindStringSubmatch(body); match != nil {
		className := match[1]
		return normalizeClassName(className)
	}

	setReturnRe := regexp.MustCompile(`SET_RETURN\s*\(\s*mrbc_instance_new\s*\([^,]+,\s*mrbc_class_(\w+)`)
	if match := setReturnRe.FindStringSubmatch(body); match != nil {
		className := match[1]
		return normalizeClassName(className)
	}

	if strings.Contains(body, "return self") || strings.Contains(body, "return mrb_obj_value(") {
		return "Self"
	}

	if strings.HasSuffix(methodName, "?") {
		return "Bool"
	}

	return "Untyped"
}

func normalizeClassName(name string) string {
	if len(name) == 0 {
		return name
	}
	if len(name) <= 2 {
		return strings.ToUpper(name)
	}
	return strings.ToUpper(name[:1]) + name[1:]
}

func inferArguments(body string, argsSpec string) []TiArgument {
	arguments := []TiArgument{}

	if strings.Contains(argsSpec, "MRB_ARGS_NONE()") {
		return arguments
	}

	if strings.Contains(argsSpec, "MRB_ARGS_ANY()") {
		arguments = append(arguments, TiArgument{
			Type: []string{"Untyped"},
			Key:  "*args",
		})
		return arguments
	}

	numRequired := 0
	numOptional := 0
	hasRest := false
	numPost := 0
	hasBlock := false

	argsReqRe := regexp.MustCompile(`MRB_ARGS_REQ\((\d+)\)`)
	if match := argsReqRe.FindStringSubmatch(argsSpec); match != nil {
		fmt.Sscanf(match[1], "%d", &numRequired)
	}

	argsOptRe := regexp.MustCompile(`MRB_ARGS_OPT\((\d+)\)`)
	if match := argsOptRe.FindStringSubmatch(argsSpec); match != nil {
		fmt.Sscanf(match[1], "%d", &numOptional)
	}

	if strings.Contains(argsSpec, "MRB_ARGS_REST()") {
		hasRest = true
	}

	argsPostRe := regexp.MustCompile(`MRB_ARGS_POST\((\d+)\)`)
	if match := argsPostRe.FindStringSubmatch(argsSpec); match != nil {
		fmt.Sscanf(match[1], "%d", &numPost)
	}

	if strings.Contains(argsSpec, "MRB_ARGS_BLOCK()") {
		hasBlock = true
	}

	getArgsRe := regexp.MustCompile(`mrb_get_args\s*\(\s*\w+\s*,\s*"([^"]+)"`)
	if match := getArgsRe.FindStringSubmatch(body); match != nil {
		formatStr := match[1]
		inOptional := false

		for _, ch := range formatStr {
			switch ch {
			case 'i':
				argType := "Int"
				if inOptional {
					argType = "DefaultInt"
				}
				arguments = append(arguments, TiArgument{Type: []string{argType}})
			case 'f':
				argType := "Float"
				if inOptional {
					argType = "DefaultFloat"
				}
				arguments = append(arguments, TiArgument{Type: []string{argType}})
			case 's', 'z':
				argType := "String"
				if inOptional {
					argType = "DefaultString"
				}
				arguments = append(arguments, TiArgument{Type: []string{argType}})
			case 'S':
				argType := "String"
				if inOptional {
					argType = "DefaultString"
				}
				arguments = append(arguments, TiArgument{Type: []string{argType}})
			case 'A':
				argType := "Array"
				if inOptional {
					argType = "DefaultUntyped"
				}
				arguments = append(arguments, TiArgument{Type: []string{argType}})
			case 'H':
				argType := "Hash"
				if inOptional {
					argType = "DefaultUntyped"
				}
				arguments = append(arguments, TiArgument{Type: []string{argType}})
			case 'a':
				argType := "Array"
				if inOptional {
					argType = "DefaultUntyped"
				}
				arguments = append(arguments, TiArgument{Type: []string{argType}})
			case 'b':
				argType := "Bool"
				if inOptional {
					argType = "DefaultUntyped"
				}
				arguments = append(arguments, TiArgument{Type: []string{argType}})
			case 'n':
				argType := "Symbol"
				if inOptional {
					argType = "DefaultUntyped"
				}
				arguments = append(arguments, TiArgument{Type: []string{argType}})
			case 'C', 'o':
				argType := "Untyped"
				if inOptional {
					argType = "DefaultUntyped"
				}
				arguments = append(arguments, TiArgument{Type: []string{argType}})
			case '|':
				inOptional = true
			case '*':
				arguments = append(arguments, TiArgument{Type: []string{"Untyped"}, Key: "*args"})
			case '&':
				arguments = append(arguments, TiArgument{Type: []string{"DefaultBlock"}})
			case '!', '?':
			}
		}
		return arguments
	}

	if numRequired == 0 && numOptional == 0 && !hasRest && numPost == 0 && !hasBlock {
		getArgTypeRe := regexp.MustCompile(`GET_(\w+)_ARG\s*\(\s*(\d+)\s*\)`)
		getArgMatches := getArgTypeRe.FindAllStringSubmatch(body, -1)
		argTypes := make(map[int]string)

		for _, match := range getArgMatches {
			argType := match[1]
			argIndex := 0
			fmt.Sscanf(match[2], "%d", &argIndex)

			switch argType {
			case "INT":
				argTypes[argIndex] = "Int"
			case "FLOAT":
				argTypes[argIndex] = "Float"
			case "STRING":
				argTypes[argIndex] = "String"
			default:
				argTypes[argIndex] = "Untyped"
			}
		}

		argCheckRe := regexp.MustCompile(`if\s*\(\s*argc\s*>=\s*(\d+)\s*\)`)
		argCheckMatches := argCheckRe.FindAllStringSubmatch(body, -1)
		maxArgc := 0
		for _, match := range argCheckMatches {
			argc := 0
			fmt.Sscanf(match[1], "%d", &argc)
			if argc > maxArgc {
				maxArgc = argc
			}
		}

		if len(argTypes) > 0 {
			maxIdx := 0
			for idx := range argTypes {
				if idx > maxIdx {
					maxIdx = idx
				}
			}

			for i := 1; i <= maxIdx; i++ {
				argType := "Untyped"
				if t, ok := argTypes[i]; ok {
					argType = t
				}

				if maxArgc > 0 && i >= maxArgc {
					if argType == "Untyped" {
						argType = "DefaultUntyped"
					} else {
						argType = "Default" + argType
					}
				}

				arguments = append(arguments, TiArgument{
					Type: []string{argType},
				})
			}
			return arguments
		}
	}

	for i := 0; i < numRequired; i++ {
		arguments = append(arguments, TiArgument{
			Type: []string{"Untyped"},
		})
	}

	for i := 0; i < numOptional; i++ {
		arguments = append(arguments, TiArgument{
			Type: []string{"DefaultUntyped"},
		})
	}

	if hasRest {
		arguments = append(arguments, TiArgument{
			Type: []string{"Untyped"},
			Key:  "*args",
		})
	}

	for i := 0; i < numPost; i++ {
		arguments = append(arguments, TiArgument{
			Type: []string{"Untyped"},
		})
	}

	if hasBlock {
		arguments = append(arguments, TiArgument{
			Type: []string{"DefaultBlock"},
		})
	}

	return arguments
}
