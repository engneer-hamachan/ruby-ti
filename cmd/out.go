package cmd

import (
	"fmt"
	"slices"
	"sort"
	"strconv"
	"ti/base"
	"ti/builtin"
	"ti/parser"
	"unicode"
)

const (
	prefixDefinitionTarget = "@"
	prefixSignature        = "%"
	prefixInheritance      = "$"
	separator              = ":::"
)

func PrintDefineInfosForPlugin(infos []string) {
	for _, info := range infos {
		fmt.Println(info)
	}
}

func PrintAllErrorsForPlugin(p parser.Parser) {
	for _, err := range p.Errors {
		fmt.Println(err)
	}
}

func PrintAllDefinitionsForLsp(p parser.Parser) {
	printDefinitionTarget(
		p.LspSudjestTargetT.DefinedFrame,
		p.LspSudjestTargetT.DefinedClass,
	)

	printMatchingSignatures(p)
	printInheritanceMap()
}

func PrintSuggestionsForLsp(p parser.Parser) {
	sortedSignatures := base.GetSortedTSignatures()

	for _, sig := range sortedSignatures {
		objectClass := p.LspSudjestTargetT.GetObjectClass()
		if objectClass == "Identifier" {
			objectClass = ""
		}

		if objectClass == "" && slices.Contains([]string{"", "Kernel"}, sig.Class) {
			printSuggestion(sig.Method, sig.Detail)
			continue
		}

		if isSuggest(p, objectClass, sig) {
			printSuggestion(sig.Method, sig.Detail)
		}
	}
}

func printDefinitionTarget(frame, class string) {
	fmt.Println(prefixDefinitionTarget + frame + separator + class)
}

func printMatchingSignatures(p parser.Parser) {
	for _, sig := range base.TSignatures {
		if p.LspSudjestTargetT.IsStatic == sig.IsStatic {
			printSignature(sig)
		}
	}
}

func printInheritanceMap() {
	for classNode, parents := range base.ClassInheritanceMap {
		for _, parent := range parents {
			printInheritance(classNode, parent)
		}
	}
}

func printSignature(sig base.Sig) {
	line := prefixSignature +
		sig.Frame + separator +
		sig.Class + separator +
		sig.Method + separator +
		sig.FileName + separator +
		strconv.Itoa(sig.Row)

	fmt.Println(line)
}

func printInheritance(child, parent base.ClassNode) {
	line := prefixInheritance +
		child.Frame + separator +
		child.Class + separator +
		parent.Frame + separator +
		parent.Class

	fmt.Println(line)
}

func printSuggestion(contents, detail string) {
	fmt.Println(prefixSignature + contents + separator + detail)
}

func isSuggest(p parser.Parser, objectClass string, sig base.Sig) bool {
	tmp := p.LspSudjestTargetT.GetBeforeEvaluateCode()

	switch tmp {
	case "Integer":
		tmp = "integer"
	case "Float":
		tmp = "float"
	case "Unknown":
		tmp = "unknown"
	}

	if len(tmp) < 1 {
		return false
	}

	if unicode.IsUpper(rune(tmp[0])) != sig.IsStatic {
		return false
	}

	if sig.Class == "" {
		return false
	}

	if sig.Class == "Kernel" {
		return false
	}

	if sig.Class == objectClass {
		return true
	}

	objectFrame := p.LspSudjestTargetT.GetFrame()

	if objectFrame == "" && slices.Contains(base.BuiltinClasses, objectClass) {
		objectFrame = "Builtin"
	}

	return isParentClass(sig, objectFrame, objectClass)
}

func isParentClass(sig base.Sig, frame, class string) bool {
	if sig.Method == "new" {
		return false
	}

	if sig.Frame == frame && sig.Class == class {
		return true
	}

	classNode := base.ClassNode{Frame: frame, Class: class}

	for _, parentNode := range base.ClassInheritanceMap[classNode] {
		if parentNode.Class == sig.Frame && parentNode.Frame == sig.Class {
			return true
		}

		if isParentClass(sig, parentNode.Frame, parentNode.Class) {
			return true
		}
	}

	return false
}

func PrintAllTypes() {
	sorted := make([]string, len(builtin.AllTypeNames))
	copy(sorted, builtin.AllTypeNames)
	sort.Strings(sorted)

	for _, name := range sorted {
		fmt.Println(name)
	}
}
