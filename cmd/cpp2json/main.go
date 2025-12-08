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

type JSONConfig struct {
	Frame          string   `json:"frame"`
	Class          string   `json:"class"`
	Extends        []string `json:"extends,omitempty"`
	Constants      []any    `json:"constants,omitempty"`
	ClassMethods   []Method `json:"class_methods,omitempty"`
	InstanceMethods []Method `json:"instance_methods,omitempty"`
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

	for _, method := range methods {
		funcInfo, found := functions[method.FuncName]
		if !found {
			continue
		}

		methodDef := analyzeFunction(funcInfo, method.MethodName)

		if isModule || method.MethodType == "class" {
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
}

func extractMethods(content string) []MethodDefinition {
	var methods []MethodDefinition

	defineRe := regexp.MustCompile(`mrbc_define_(method|class_method)\s*\(\s*\w+\s*,\s*\w+\s*,\s*"([^"]+)"\s*,\s*(\w+)\s*\)`)
	matches := defineRe.FindAllStringSubmatch(content, -1)

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

	return methods
}

func extractFunctions(content string) map[string]FunctionInfo {
	functions := make(map[string]FunctionInfo)

	funcRe := regexp.MustCompile(`(?s)void\s+(\w+)\s*\([^)]*\)\s*\{(.*?)\n\}`)
	matches := funcRe.FindAllStringSubmatch(content, -1)

	for _, match := range matches {
		funcName := match[1]
		funcBody := match[2]
		functions[funcName] = FunctionInfo{
			Name: funcName,
			Body: funcBody,
		}
	}

	return functions
}

func analyzeFunction(funcInfo FunctionInfo, methodName string) Method {
	method := Method{
		Name:      methodName,
		Arguments: []Argument{},
	}

	returnType := inferReturnType(funcInfo.Body, methodName)
	method.ReturnType = ReturnType{Type: []string{returnType}}

	arguments := inferArguments(funcInfo.Body)
	method.Arguments = arguments

	document := extractDocument(funcInfo.Body)
	if document != "" {
		method.Document = document
	}

	return method
}

func inferReturnType(body string, methodName string) string {
	if strings.Contains(body, "SET_NIL_RETURN()") {
		return "Nil"
	}
	if strings.Contains(body, "SET_INT_RETURN(") {
		return "Int"
	}
	if strings.Contains(body, "SET_FLOAT_RETURN(") {
		return "Float"
	}
	if strings.Contains(body, "SET_TRUE_RETURN()") || strings.Contains(body, "SET_FALSE_RETURN()") || strings.Contains(body, "SET_BOOL_RETURN(") {
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

func inferArguments(body string) []Argument {
	arguments := []Argument{}

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

	for i := 1; i <= maxArgc; i++ {
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

	return arguments
}

func extractDocument(body string) string {
	docRe := regexp.MustCompile(`/\*[^*]*\*\s*Method:\s*[^*]+\s*\*([^*]+)\*`)
	if match := docRe.FindStringSubmatch(body); match != nil {
		return strings.TrimSpace(match[1])
	}

	lines := strings.Split(body, "\n")
	for _, line := range lines {
		if strings.Contains(line, "//") {
			comment := strings.TrimSpace(strings.SplitN(line, "//", 2)[1])
			if comment != "" && !strings.HasPrefix(comment, "TODO") {
				return comment
			}
		}
	}

	return ""
}
