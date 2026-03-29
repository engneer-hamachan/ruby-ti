package main

import (
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"unicode"
)

//go:embed rbs_ast.rb
var rbsAstScript string

// --- RBS AST structures ---

type RBSDeclaration struct {
	Declaration string         `json:"declaration"`
	Name        string         `json:"name"`
	TypeParams  []any          `json:"type_params"`
	Members     []RBSMember    `json:"members"`
	SuperClass  *RBSSuperClass `json:"super_class"`
	Comment     *RBSComment    `json:"comment"`
}

type RBSSuperClass struct {
	Name string `json:"name"`
	Args []any  `json:"args"`
}

type RBSComment struct {
	String string `json:"string"`
}

type RBSMember struct {
	Member      string        `json:"member"`
	Declaration string        `json:"declaration"`
	Name        string        `json:"name"`
	Kind        string        `json:"kind"`
	Visibility  string        `json:"visibility"`
	Overloads   []RBSOverload `json:"overloads"`
	Comment     *RBSComment   `json:"comment"`

	NewName string `json:"new_name"`
	OldName string `json:"old_name"`

	Type     *RBSType `json:"type"`
	IvarName *string  `json:"ivar_name"`

	Members    []RBSMember    `json:"members"`
	SuperClass *RBSSuperClass `json:"super_class"`
	TypeParams []any          `json:"type_params"`
}

type RBSOverload struct {
	MethodType RBSMethodType `json:"method_type"`
}

type RBSMethodType struct {
	TypeParams []any       `json:"type_params"`
	Type       RBSFuncType `json:"type"`
	Block      *RBSBlock   `json:"block"`
}

type RBSFuncType struct {
	RequiredPositionals []RBSParam          `json:"required_positionals"`
	OptionalPositionals []RBSParam          `json:"optional_positionals"`
	RestPositionals     *RBSParam           `json:"rest_positionals"`
	TrailingPositionals []RBSParam          `json:"trailing_positionals"`
	RequiredKeywords    map[string]RBSParam `json:"required_keywords"`
	OptionalKeywords    map[string]RBSParam `json:"optional_keywords"`
	RestKeywords        *RBSParam           `json:"rest_keywords"`
	ReturnType          RBSType             `json:"return_type"`
}

type RBSParam struct {
	Type *RBSType `json:"type"`
	Name string   `json:"name"`
}

type RBSType struct {
	Class   string    `json:"class"`
	Name    string    `json:"name"`
	Args    []RBSType `json:"args"`
	Type    *RBSType  `json:"type"`
	Types   []RBSType `json:"types"`
	Literal string    `json:"literal"`
}

type RBSBlock struct {
	Type     RBSFuncType `json:"type"`
	Required bool        `json:"required"`
}

// --- ti-config output structures ---

type TiArgument struct {
	Type       []string `json:"type,omitempty"`
	Key        string   `json:"key,omitempty"`
	IsAsterisk bool     `json:"is_asterisk,omitempty"`
	IsDefault  bool     `json:"is_default,omitempty"`
}

type TiReturnType struct {
	Type []string `json:"type"`
}

type TiMethod struct {
	Name            string       `json:"name"`
	Arguments       []TiArgument `json:"arguments"`
	ReturnType      TiReturnType `json:"return_type"`
	BlockParameters []string     `json:"block_parameters,omitempty"`
	Document        string       `json:"document,omitempty"`
}

type TiConstant struct {
	Name       string       `json:"name"`
	ReturnType TiReturnType `json:"return_type"`
}

type TiProperty struct {
	Name   string   `json:"name"`
	Type   []string `json:"type"`
	Access string   `json:"access"`
}

type TiInstanceVar struct {
	Name string   `json:"name"`
	Type []string `json:"type"`
}

type TiClassConfig struct {
	Frame              string          `json:"frame"`
	Class              string          `json:"class"`
	Type               string          `json:"type,omitempty"`
	Extends            []string        `json:"extends"`
	InstanceMethods    []TiMethod      `json:"instance_methods"`
	ClassMethods       []TiMethod      `json:"class_methods,omitempty"`
	Constants          []TiConstant    `json:"constants,omitempty"`
	InstanceProperties []TiProperty    `json:"instance_properties,omitempty"`
	InstanceVariables  []TiInstanceVar `json:"instance_variables,omitempty"`
}

// --- type alias ---

type typeAliasMap map[string]RBSType

func collectTypeAliases(members []RBSMember) typeAliasMap {
	aliases := make(typeAliasMap)
	for _, m := range members {
		if m.Declaration == "alias" && m.Member == "" && m.Type != nil {
			aliases[m.Name] = *m.Type
		}
	}
	return aliases
}

// --- type conversion ---

func convertType(t RBSType, aliases typeAliasMap, className string) []string {
	switch t.Class {
	case "class_instance":
		return convertClassInstance(t, aliases, className)
	case "bool":
		return []string{"Bool"}
	case "nil":
		return []string{"NilClass"}
	case "void":
		return []string{"NilClass"}
	case "untyped":
		return []string{"Untyped"}
	case "self":
		return []string{"Self"}
	case "instance":
		return []string{className}
	case "variable":
		return []string{"Untyped"}
	case "optional":
		if t.Type != nil {
			inner := convertType(*t.Type, aliases, className)
			if !containsType(inner, "NilClass") {
				inner = append(inner, "NilClass")
			}
			return inner
		}
		return []string{"NilClass"}
	case "union":
		return convertUnion(t, aliases, className)
	case "literal":
		return convertLiteral(t)
	case "tuple":
		return []string{"Array"}
	case "alias":
		// RBS built-in type aliases
		switch t.Name {
		case "int":
			return []string{"Int"}
		case "float":
			return []string{"Float"}
		case "string", "path", "encoding":
			return []string{"String"}
		case "boolish", "bool":
			return []string{"Bool"}
		case "real":
			return []string{"Float", "Int"}
		case "interned":
			return []string{"Symbol", "String"}
		case "io":
			return []string{"Untyped"}
		case "range":
			return []string{"Range"}
		case "array":
			return []string{"Array"}
		case "hash":
			return []string{"Hash"}
		}
		if resolved, ok := aliases[t.Name]; ok {
			return convertType(resolved, aliases, className)
		}
		return []string{"Untyped"}
	case "intersection":
		if len(t.Types) > 0 {
			return convertType(t.Types[0], aliases, className)
		}
		return []string{"Untyped"}
	default:
		return []string{"Untyped"}
	}
}

func convertClassInstance(t RBSType, aliases typeAliasMap, className string) []string {
	switch t.Name {
	case "Integer", "int":
		return []string{"Int"}
	case "Float":
		return []string{"Float"}
	case "String":
		return []string{"String"}
	case "Symbol":
		return []string{"Symbol"}
	case "NilClass":
		return []string{"NilClass"}
	case "TrueClass", "FalseClass":
		return []string{"Bool"}
	case "Object":
		return []string{"Untyped"}
	case "Array":
		if len(t.Args) > 0 {
			inner := convertType(t.Args[0], aliases, className)
			if len(inner) == 1 {
				return []string{"[" + inner[0] + "]"}
			}
		}
		return []string{"Array"}
	case "Hash":
		return []string{"Hash"}
	default:
		return []string{t.Name}
	}
}

func convertUnion(t RBSType, aliases typeAliasMap, className string) []string {
	var result []string
	seen := make(map[string]bool)
	for _, ut := range t.Types {
		types := convertType(ut, aliases, className)
		for _, typ := range types {
			if isSymbolLiteral(ut) {
				typ = "Symbol"
			}
			if !seen[typ] {
				seen[typ] = true
				result = append(result, typ)
			}
		}
	}
	return result
}

func isSymbolLiteral(t RBSType) bool {
	return t.Class == "literal" && strings.HasPrefix(t.Literal, ":")
}

func convertLiteral(t RBSType) []string {
	lit := t.Literal
	if strings.HasPrefix(lit, ":") {
		return []string{"Symbol"}
	}
	if lit == "true" || lit == "false" {
		return []string{"Bool"}
	}
	if len(lit) > 0 && (lit[0] >= '0' && lit[0] <= '9' || lit[0] == '-') {
		if strings.Contains(lit, ".") {
			return []string{"Float"}
		}
		return []string{"Int"}
	}
	if strings.HasPrefix(lit, "\"") || strings.HasPrefix(lit, "'") {
		return []string{"String"}
	}
	return []string{"Untyped"}
}

func containsType(types []string, target string) bool {
	for _, t := range types {
		if t == target {
			return true
		}
	}
	return false
}

// --- declaration conversion ---

func convertDeclarations(decls []RBSDeclaration, parentName string) []TiClassConfig {
	var configs []TiClassConfig

	for _, decl := range decls {
		className := decl.Name
		if parentName != "" {
			className = parentName + "::" + decl.Name
		}

		aliases := collectTypeAliases(decl.Members)

		config := TiClassConfig{
			Frame:           "Builtin",
			Class:           className,
			Extends:         []string{},
			InstanceMethods: []TiMethod{},
		}

		if decl.Declaration == "module" {
			config.Type = "module"
		}

		if decl.SuperClass != nil && decl.SuperClass.Name != "" {
			superName := decl.SuperClass.Name
			if parentName != "" && !strings.Contains(superName, "::") {
				superName = parentName + "::" + superName
			}
			config.Extends = []string{superName}
		}

		instanceMethodDefs := make(map[string][]TiMethod)
		classMethodDefs := make(map[string][]TiMethod)

		for _, member := range decl.Members {
			memberType := member.Member
			if memberType == "" {
				memberType = member.Declaration
			}

			switch memberType {
			case "method_definition":
				if member.Visibility == "private" {
					continue
				}
				methods := convertMethodDefinition(member, aliases, className)
				if member.Kind == "singleton" {
					config.ClassMethods = append(config.ClassMethods, methods...)
					classMethodDefs[member.Name] = methods
				} else {
					config.InstanceMethods = append(config.InstanceMethods, methods...)
					instanceMethodDefs[member.Name] = methods
				}

			case "alias":
				if member.Member == "alias" {
					oldName := member.OldName
					newName := member.NewName
					if member.Kind == "singleton" {
						if orig, ok := classMethodDefs[oldName]; ok {
							for _, m := range orig {
								aliased := m
								aliased.Name = newName
								config.ClassMethods = append(config.ClassMethods, aliased)
							}
							classMethodDefs[newName] = classMethodDefs[oldName]
						}
					} else {
						if orig, ok := instanceMethodDefs[oldName]; ok {
							for _, m := range orig {
								aliased := m
								aliased.Name = newName
								config.InstanceMethods = append(config.InstanceMethods, aliased)
							}
							instanceMethodDefs[newName] = instanceMethodDefs[oldName]
						}
					}
				}

			case "attr_reader":
				prop := convertAttrToProperty(member, "reader", aliases, className)
				config.InstanceProperties = append(config.InstanceProperties, prop)

			case "attr_accessor":
				prop := convertAttrToProperty(member, "accessor", aliases, className)
				config.InstanceProperties = append(config.InstanceProperties, prop)

			case "instance_variable":
				ivar := convertInstanceVariable(member, aliases, className)
				config.InstanceVariables = append(config.InstanceVariables, ivar)

			case "class", "module":
				nestedDecl := memberToDeclaration(member)
				nested := convertDeclarations([]RBSDeclaration{nestedDecl}, className)
				configs = append(configs, nested...)

			case "constant":
				if member.Type != nil {
					constant := TiConstant{
						Name: member.Name,
						ReturnType: TiReturnType{
							Type: convertType(*member.Type, aliases, className),
						},
					}
					config.Constants = append(config.Constants, constant)
				}
			}
		}

		configs = append(configs, config)
	}

	return configs
}

func convertMethodDefinition(member RBSMember, aliases typeAliasMap, className string) []TiMethod {
	var methods []TiMethod

	doc := ""
	if member.Comment != nil && member.Comment.String != "" {
		doc = strings.TrimSpace(member.Comment.String)
	}

	for _, overload := range member.Overloads {
		method := TiMethod{
			Name:      member.Name,
			Arguments: convertArguments(overload.MethodType.Type, aliases, className),
			ReturnType: TiReturnType{
				Type: convertType(overload.MethodType.Type.ReturnType, aliases, className),
			},
		}

		if overload.MethodType.Block != nil {
			var blockParams []string
			for _, param := range overload.MethodType.Block.Type.RequiredPositionals {
				if param.Type != nil {
					types := convertType(*param.Type, aliases, className)
					blockParams = append(blockParams, types...)
				}
			}
			if len(blockParams) > 0 {
				method.BlockParameters = blockParams
			}
		}

		if doc != "" {
			method.Document = doc
		}

		methods = append(methods, method)
	}

	return methods
}

func convertArguments(funcType RBSFuncType, aliases typeAliasMap, className string) []TiArgument {
	var args []TiArgument

	for _, param := range funcType.RequiredPositionals {
		if param.Type != nil {
			types := convertType(*param.Type, aliases, className)
			args = append(args, TiArgument{Type: types})
		}
	}

	for _, param := range funcType.OptionalPositionals {
		if param.Type != nil {
			types := convertType(*param.Type, aliases, className)
			args = append(args, TiArgument{Type: types, IsDefault: true})
		}
	}

	if funcType.RestPositionals != nil {
		args = append(args, TiArgument{IsAsterisk: true})
	}

	for _, param := range funcType.TrailingPositionals {
		if param.Type != nil {
			types := convertType(*param.Type, aliases, className)
			args = append(args, TiArgument{Type: types})
		}
	}

	for name, kw := range funcType.RequiredKeywords {
		if kw.Type != nil {
			types := convertType(*kw.Type, aliases, className)
			args = append(args, TiArgument{
				Type: types,
				Key:  name + ":",
			})
		}
	}

	for name, kw := range funcType.OptionalKeywords {
		if kw.Type != nil {
			types := convertType(*kw.Type, aliases, className)
			args = append(args, TiArgument{
				Type:      types,
				Key:       name + ":",
				IsDefault: true,
			})
		}
	}

	if args == nil {
		args = []TiArgument{}
	}
	return args
}

func convertAttrToProperty(member RBSMember, access string, aliases typeAliasMap, className string) TiProperty {
	var types []string
	if member.Type != nil {
		types = convertType(*member.Type, aliases, className)
	} else {
		types = []string{"Untyped"}
	}
	return TiProperty{
		Name:   member.Name,
		Type:   types,
		Access: access,
	}
}

func convertInstanceVariable(member RBSMember, aliases typeAliasMap, className string) TiInstanceVar {
	name := strings.TrimPrefix(member.Name, "@")
	var types []string
	if member.Type != nil {
		types = convertType(*member.Type, aliases, className)
	} else {
		types = []string{"Untyped"}
	}
	return TiInstanceVar{
		Name: name,
		Type: types,
	}
}

func memberToDeclaration(member RBSMember) RBSDeclaration {
	return RBSDeclaration{
		Declaration: member.Declaration,
		Name:        member.Name,
		Members:     member.Members,
		SuperClass:  member.SuperClass,
		TypeParams:  member.TypeParams,
		Comment:     member.Comment,
	}
}

// --- CLI ---

func main() {
	output := flag.String("o", "", "output path (directory/ or file)")
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Fprintf(os.Stderr, "Usage: rbs2json [options] <input.rbs>\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	inputFile := flag.Arg(0)

	astJSON, err := execRubyScript(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	var decls []RBSDeclaration
	if err := json.Unmarshal(astJSON, &decls); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing AST JSON: %v\n", err)
		os.Exit(1)
	}

	configs := convertDeclarations(decls, "")

	if err := writeOutput(configs, *output); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func execRubyScript(inputFile string) ([]byte, error) {
	tmpFile, err := os.CreateTemp("", "rbs_ast_*.rb")
	if err != nil {
		return nil, err
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.WriteString(rbsAstScript); err != nil {
		tmpFile.Close()
		return nil, err
	}
	tmpFile.Close()

	cmd := exec.Command("ruby", tmpFile.Name(), inputFile)
	cmd.Stderr = os.Stderr
	return cmd.Output()
}

func classNameToFileName(className string) string {
	parts := strings.Split(className, "::")
	var snakeParts []string
	for _, part := range parts {
		snakeParts = append(snakeParts, toSnakeCase(part))
	}
	return strings.Join(snakeParts, "_") + ".json"
}

func toSnakeCase(s string) string {
	var result []rune
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				result = append(result, '_')
			}
			result = append(result, unicode.ToLower(r))
		} else {
			result = append(result, r)
		}
	}
	return string(result)
}

func writeOutput(configs []TiClassConfig, outputPath string) error {
	if outputPath == "" {
		var data any
		if len(configs) == 1 {
			data = configs[0]
		} else {
			data = configs
		}
		jsonData, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			return err
		}
		fmt.Println(string(jsonData))
		return nil
	}

	info, err := os.Stat(outputPath)
	isDir := err == nil && info.IsDir()

	if !isDir && !strings.HasSuffix(outputPath, "/") {
		if len(configs) != 1 {
			return fmt.Errorf("multiple classes (%d), use directory output", len(configs))
		}
		return writeJSONFile(outputPath, configs[0])
	}

	os.MkdirAll(outputPath, 0755)
	for _, config := range configs {
		fileName := classNameToFileName(config.Class)
		filePath := filepath.Join(outputPath, fileName)
		if err := writeJSONFile(filePath, config); err != nil {
			return err
		}
		fmt.Fprintf(os.Stderr, "Generated: %s\n", filePath)
	}
	return nil
}

func writeJSONFile(path string, data any) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, append(jsonData, '\n'), 0644)
}
