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
	targetT := p.LspSudjestTargetT
	isPrinted := false

	for _, sig := range base.GetSortedTSignatures() {
		if isSuggestForKernelOrObjectClass(targetT, sig.Class) {
			printSuggestion(sig.Method, sig.Detail)
			isPrinted = true
			continue
		}

		if isSuggest(targetT, sig) {
			printSuggestion(sig.Method, sig.Detail)
			isPrinted = true
		}
	}

	if isPrinted {
		return
	}

	if targetT.IsIdentifierType() && unicode.IsUpper(rune(targetT.ToString()[0])) {
		printAllClasses()
	}
}

func printAllClasses() {
	classSet := make(map[string]bool)

	for _, sig := range base.TSignatures {
		if sig.Class != "" && sig.Class != "Kernel" {
			switch sig.Frame {
			case "", "Builtin":
				classSet[sig.Class] = true
			default:
				classSet[sig.Frame+"::"+sig.Class] = true
			}
		}
	}

	classes := make([]string, 0, len(classSet))
	for className := range classSet {
		classes = append(classes, className)
	}

	sort.Strings(classes)

	for _, className := range classes {
		fmt.Println(prefixSignature + className + separator + className)
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

func isSuggestForKernelOrObjectClass(targetT base.T, sigClass string) bool {
	if unicode.IsUpper(rune(targetT.ToString()[0])) {
		return false
	}

	return slices.Contains([]string{"", "Kernel"}, sigClass)
}

func calculateObjectClassAndIsStatic(targetT base.T) (string, bool) {
	target := targetT.ToString()
	beforeCode := targetT.GetBeforeEvaluateCode()
	isStaticTarget := unicode.IsUpper(rune(target[0]))

	if targetT.GetType() == base.SELF && !isStaticTarget {
		return targetT.DefinedClass, false
	}

	// 1, '1', 1.1, [], {} and more...
	switch targetT.GetType() {
	case base.INT, base.FLOAT:
		return beforeCode, false

	case base.ARRAY, base.HASH, base.STRING, base.OBJECT:
		return targetT.GetObjectClass(), false

	case base.UNKNOWN:
		// static top level method in class
		if !isStaticTarget {
			if targetT.DefinedMethod == "" {
				return targetT.DefinedClass, true
			}

			return targetT.DefinedClass, targetT.IsStatic
		}
	}

	if beforeCode == "" || isStaticTarget {
		return target, isStaticTarget
	}

	return targetT.GetObjectClass(), isStaticTarget
}

func isSuggest(targetT base.T, sig base.Sig) bool {
	if sig.Class == "" {
		return false
	}

	if sig.Class == "Kernel" {
		return false
	}

	objectClass, isStaticTarget := calculateObjectClassAndIsStatic(targetT)

	if len(objectClass) < 1 {
		return false
	}

	if isStaticTarget != sig.IsStatic {
		return false
	}

	if sig.Class == objectClass {
		return true
	}

	return isParentClass(sig, targetT.GetFrame(), objectClass)
}

func isParentClass(sig base.Sig, frame, class string) bool {
	if sig.Method == "new" {
		return false
	}

	if frame == "" && slices.Contains(base.BuiltinClasses, class) {
		frame = "Builtin"
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

func PrintTargetClassExtends() {
	className := getTargetClass()

	for classNode, parents := range base.ClassInheritanceMap {
		if classNode.Class == className {
			for _, parent := range parents {
				switch parent.Class {
				case "":
					fmt.Println("Object")
				default:
					fmt.Println(parent.Class)
				}
			}

			return
		}
	}
}
