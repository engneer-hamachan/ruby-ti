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

func PrintDefineInfos(infos []string) {
	for _, info := range infos {
		fmt.Println(info)
	}
}

func PrintDefinitionTarget(frame, class string) {
	fmt.Println(prefixDefinitionTarget + frame + separator + class)
}

func PrintMatchingSignatures(p parser.Parser) {
	for _, sig := range base.TSignatures {
		if p.LspSudjestTargetT.IsStatic == sig.IsStatic {
			printSignature(sig)
		}
	}
}

func PrintInheritanceMap() {
	for classNode, parents := range base.ClassInheritanceMap {
		for _, parent := range parents {
			printInheritance(classNode, parent)
		}
	}
}

func PrintLspSuggestions(p parser.Parser) {
	for _, sig := range base.TSignatures {
		objectClass := p.LspSudjestTargetT.GetObjectClass()
		if objectClass == "Identifier" {
			objectClass = ""
		}

		if objectClass == "" && slices.Contains([]string{"", "Kernel"}, sig.Class) {
			printSuggestion(sig.Contents, sig.Detail)
			continue
		}

		if isSudjest(objectClass, sig) {
			tmp := p.LspSudjestTargetT.GetBeforeEvaluateCode()
			if len(tmp) > 0 && unicode.IsUpper(rune(tmp[0])) == sig.IsStatic {
				printSuggestion(sig.Contents, sig.Detail)
			}
		}
	}
}

func printSignature(sig base.Sig) {
	line := prefixSignature +
		sig.Frame + separator +
		sig.Class + separator +
		sig.Contents + separator +
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

func isSudjest(objectClass string, sig base.Sig) bool {
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
