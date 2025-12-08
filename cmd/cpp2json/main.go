package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Argument struct {
	Type []string `json:"type"`
	Key  string   `json:"key,omitempty"`
}

type ReturnType struct {
	Type []string `json:"type"`
}

type Method struct {
	Name       string     `json:"name"`
	Arguments  []Argument `json:"arguments"`
	ReturnType ReturnType `json:"return_type"`
	Document   string     `json:"document,omitempty"`
}

type Constant struct {
	Name       string     `json:"name"`
	ReturnType ReturnType `json:"return_type"`
}

type JSONConfig struct {
	Frame           string     `json:"frame"`
	Class           string     `json:"class"`
	Extends         []string   `json:"extends,omitempty"`
	Constants       []Constant `json:"constants,omitempty"`
	ClassMethods    []Method   `json:"class_methods,omitempty"`
	InstanceMethods []Method   `json:"instance_methods,omitempty"`
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

func parseFile(content string, className string, isModule bool) JSONConfig {
	config := JSONConfig{
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

	mrbDefineRe := regexp.MustCompile(`mrb_define_(class_)?method_id\s*\(\s*\w+\s*,\s*\w+\s*,\s*MRB_SYM(_Q)?\((\w+)\)\s*,\s*(\w+)\s*,\s*([^)]+\))`)
	mrbMatches := mrbDefineRe.FindAllStringSubmatch(content, -1)

	for _, match := range mrbMatches {
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

	return methods
}

func extractConstants(content string) []Constant {
	var constants []Constant

	constRe := regexp.MustCompile(`mrb_define_const_id\s*\(\s*\w+\s*,\s*\w+\s*,\s*MRB_SYM\((\w+)\)\s*,\s*mrb_(\w+)_value\(`)
	matches := constRe.FindAllStringSubmatch(content, -1)

	for _, match := range matches {
		constName := match[1]
		valueType := match[2]

		typeName := "Untyped"
		switch valueType {
		case "fixnum":
			typeName = "Int"
		case "float":
			typeName = "Float"
		case "true", "false":
			typeName = "Bool"
		case "nil":
			typeName = "Nil"
		}

		constants = append(constants, Constant{
			Name: constName,
			ReturnType: ReturnType{
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

func analyzeFunction(funcInfo FunctionInfo, methodName string, argsSpec string) Method {
	method := Method{
		Name:      methodName,
		Arguments: []Argument{},
	}

	returnType := inferReturnType(funcInfo.Body, methodName)
	method.ReturnType = ReturnType{Type: []string{returnType}}

	arguments := inferArguments(funcInfo.Body, argsSpec)
	method.Arguments = arguments
	method.Document = ""

	return method
}

func inferReturnType(body string, methodName string) string {
	if strings.Contains(body, "SET_NIL_RETURN()") || strings.Contains(body, "return mrb_nil_value()") {
		return "Nil"
	}
	if strings.Contains(body, "SET_INT_RETURN(") || strings.Contains(body, "return mrb_fixnum_value(") {
		return "Int"
	}
	if strings.Contains(body, "SET_FLOAT_RETURN(") || strings.Contains(body, "return mrb_float_value(") {
		return "Float"
	}
	if strings.Contains(body, "SET_TRUE_RETURN()") || strings.Contains(body, "SET_FALSE_RETURN()") ||
	   strings.Contains(body, "SET_BOOL_RETURN(") || strings.Contains(body, "return mrb_true_value()") ||
	   strings.Contains(body, "return mrb_false_value()") {
		return "Bool"
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

func inferArguments(body string, argsSpec string) []Argument {
	arguments := []Argument{}

	if strings.Contains(argsSpec, "MRB_ARGS_NONE()") {
		return arguments
	}

	argsReqRe := regexp.MustCompile(`MRB_ARGS_REQ\((\d+)\)`)
	if match := argsReqRe.FindStringSubmatch(argsSpec); match != nil {
		numArgs := 0
		fmt.Sscanf(match[1], "%d", &numArgs)

		if strings.Contains(body, "pin_num(mrb, NULL)") {
			arguments = append(arguments, Argument{
				Type: []string{"Int", "String", "Symbol"},
			})
			numArgs--
		} else if strings.Contains(body, "pin_num(mrb, &") {
			arguments = append(arguments, Argument{
				Type: []string{"Int", "String", "Symbol"},
			})
			arguments = append(arguments, Argument{
				Type: []string{"Int"},
			})
			numArgs -= 2
		}

		getArgsRe := regexp.MustCompile(`mrb_get_args\s*\(\s*\w+\s*,\s*"([^"]+)"`)
		if match := getArgsRe.FindStringSubmatch(body); match != nil {
			formatStr := match[1]
			for _, ch := range formatStr {
				switch ch {
				case 'i':
					arguments = append(arguments, Argument{
						Type: []string{"Int"},
					})
				case 'f':
					arguments = append(arguments, Argument{
						Type: []string{"Float"},
					})
				case 's':
					arguments = append(arguments, Argument{
						Type: []string{"String"},
					})
				case 'o':
					arguments = append(arguments, Argument{
						Type: []string{"Untyped"},
					})
				}
			}
		} else {
			for i := 0; i < numArgs; i++ {
				arguments = append(arguments, Argument{
					Type: []string{"Untyped"},
				})
			}
		}
	}

	argCheckRe := regexp.MustCompile(`if\s*\(\s*argc\s*>=\s*(\d+)\s*\)`)
	maxArgc := 0
	matches := argCheckRe.FindAllStringSubmatch(body, -1)
	for _, match := range matches {
		argc := 0
		fmt.Sscanf(match[1], "%d", &argc)
		if argc > maxArgc {
			maxArgc = argc
		}
	}

	if maxArgc > 0 && len(arguments) > 0 {
		if len(arguments) >= maxArgc {
			for i := maxArgc - 1; i < len(arguments); i++ {
				if arguments[i].Type[0] != "Untyped" {
					arguments[i].Type = []string{"Default" + arguments[i].Type[0]}
				}
			}
		}
	}

	getArgIndexRe := regexp.MustCompile(`GET_\w+_ARG\s*\(\s*(\d+)\s*\)`)
	getArgTypeRe := regexp.MustCompile(`GET_(\w+)_ARG\s*\(\s*(\d+)\s*\)`)
	argTypes := make(map[int]string)

	getArgMatches := getArgTypeRe.FindAllStringSubmatch(body, -1)
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

	hasArgcCheck := maxArgc > 0
	allArgIndices := make(map[int]bool)
	argIndexMatches := getArgIndexRe.FindAllStringSubmatch(body, -1)
	for _, match := range argIndexMatches {
		idx := 0
		fmt.Sscanf(match[1], "%d", &idx)
		allArgIndices[idx] = true
	}

	if len(arguments) == 0 && len(argTypes) > 0 {
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

			if hasArgcCheck && allArgIndices[i] {
				argType = "Default" + argType
			}

			arguments = append(arguments, Argument{
				Type: []string{argType},
			})
		}
	}

	return arguments
}
