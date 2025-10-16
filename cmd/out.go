package cmd

import (
	"fmt"
	"slices"
	"strconv"
	"ti/base"
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
	printDefinitionTarget(p.LspSudjestTargetT.DefinedFrame, p.LspSudjestTargetT.DefinedClass)
	printMatchingSignatures(p)
	printInheritanceMap()
}

func PrintSuggestionsForLsp(p parser.Parser) {
	for _, sig := range base.TSignatures {
		objectClass := p.LspSudjestTargetT.GetObjectClass()
		if objectClass == "Identifier" {
			objectClass = ""
		}

		if objectClass == "" && slices.Contains([]string{"", "Kernel"}, sig.Class) {
			printSuggestion(sig.Method, sig.Detail)
			continue
		}

		if isSuggest(objectClass, sig) {
			tmp := p.LspSudjestTargetT.GetBeforeEvaluateCode()
			if len(tmp) > 0 && unicode.IsUpper(rune(tmp[0])) == sig.IsStatic {
				printSuggestion(sig.Method, sig.Detail)
			}
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

func isSuggest(objectClass string, sig base.Sig) bool {
	if sig.Class == "" {
		return false
	}

	if sig.Class == objectClass {
		return true
	}

	return isParentClass(sig.Frame, objectClass, sig.Class)
}

func isParentClass(frame, childClass, parentClass string) bool {
	classNode := base.ClassNode{Frame: frame, Class: childClass}

	for _, parentNode := range base.ClassInheritanceMap[classNode] {
		if parentNode.Class == parentClass {
			return true
		}

		// Recursively check parent's parents
		if isParentClass(parentNode.Frame, parentNode.Class, parentClass) {
			return true
		}
	}

	return false
}
