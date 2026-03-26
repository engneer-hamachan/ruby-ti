package cmd

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"ti/base"
	"ti/builtin"
	"ti/eval"
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

func PrintDefineInfosForLlm() {
	for _, sig := range base.GetSortedTSignatures() {
		if sig.Frame != "Builtin" && sig.Method != "new" {
			if sig.Document == "" {
				sig.Document = "<no document>"
			}

			fmt.Println("## " + sig.GetPrintDetail())
			endRow := findEndRow(sig.Frame, sig.Class, sig.Method)
			fmt.Println("- file: " + sig.FileName + ":" + strconv.Itoa(sig.Row) + "-" + strconv.Itoa(endRow))
			fmt.Println("- document: " + sig.Document)

			points := base.MethodCallPoint[sig.Frame+sig.Class+sig.Method]

			// print cllers start
			if len(points) == 0 {
				fmt.Println("- callers: none")
			} else {
				fmt.Println("- callers:")

				for _, point := range points {
					if point.CallerMethod != "" {
						fmt.Println("  - method: " + point.CallerMethod)
					} else {
						fmt.Println("  - method: " + "top level")
					}

					if point.CallerClass != "" {
						fmt.Println("    - class: " + point.CallerClass)
					} else {
						fmt.Println("    - class: none")
					}

					fmt.Println("    - call point: " + point.Point)
				}
			}

			fmt.Printf("  - total callers: %d\n", len(points))
			// print cllers end

			// print cllees start
			callees := base.MethodCalleePoint[sig.Frame+sig.Class+sig.Method]

			printedCallees := 0

			if len(callees) > 0 {
				fmt.Println("- callees:")
			} else {
				fmt.Println("- callees: none")
			}

			for _, cp := range callees {
				definePoint, ok := findDefinePoint(cp.CalleeFrame, cp.CalleeClass, cp.CalleeMethod)
				if !ok {
					continue
				}
				fmt.Println("  - method: " + cp.CalleeMethod)
				if cp.CalleeClass != "" {
					fmt.Println("    - class: " + cp.CalleeClass)
				} else {
					fmt.Println("    - class: none")
				}
				fmt.Println("    - define point: " + definePoint)
				printedCallees++
			}

			fmt.Printf("  - total callees: %d\n", printedCallees)

			printSpecialComments(sig.FileName, sig.Row, endRow)

			fmt.Println()
		}
	}
}

func PrintSpecialCodeCommentsForLlm() {
	for _, sig := range base.SpecialCodeComments {
		if isCommentInAnyFunction(sig) {
			continue
		}
		fmt.Println("## " + sig.FileName + ":" + strconv.Itoa(sig.Row))
		fmt.Println("- comment: " + sig.Document)

		line := readLineFromFile(sig.FileName, sig.Row)
		if line != "" {
			fmt.Println("```")
			fmt.Println(line)
			fmt.Println("```")
		}

		fmt.Println()
	}
}

func PrintAllDefinitionsForLlm() {
	class := ""
	previewClass := ""

	targetClassName := getTargetClass()

	for _, sig := range base.GetSortedTSignaturesByClass() {
		if sig.Class == "" {
			class = "Object"
		} else {
			class = sig.Class
		}

		if targetClassName != "" && targetClassName != class {
			continue
		}

		if class != previewClass {
			fmt.Println("---")
			fmt.Println("# " + class)
			previewClass = class
		}

		fmt.Println("## " + sig.Detail)
		fmt.Println("- document: " + sig.Document)

		if sig.IsStatic {
			fmt.Println("- isStatic: " + "true")
		} else {
			fmt.Println("- isStatic: " + "false")
		}
	}
}

func PrintAllClassesForLlm() {
	class := ""
	previewClass := ""

	for _, sig := range base.GetSortedTSignaturesByClass() {
		if sig.Class == "" {
			class = "Object"
		} else {
			class = sig.Class
		}

		if class != previewClass {
			fmt.Println(class)
			previewClass = class
		}
	}
}

func sigHasCallPoints(sig base.Sig) bool {
	key := sig.Frame + sig.Class + sig.Method
	return len(base.MethodCallPoint[key]) > 0 || len(base.MethodCalleePoint[key]) > 0
}

func PrintLlmNav() bool {
	target := getTarget()

	if target != "" {
		printLlmNavDetail(target)
		return true
	}

	printLlmNavList()
	return false
}

func printLlmNavList() {
	class := ""
	previewClass := ""
	topLevelMethods := []string{}

	fmt.Println("# Classes")
	for _, sig := range base.GetSortedTSignaturesByClass() {
		if sig.Frame == "Builtin" {
			continue
		}

		if !sigHasCallPoints(sig) {
			continue
		}

		if sig.Class == "" {
			topLevelMethods = append(topLevelMethods, sig.Method)
			continue
		}

		class = sig.Class

		if class != previewClass {
			fmt.Println(class)
			previewClass = class
		}
	}

	if len(topLevelMethods) > 0 {
		fmt.Println("---")
		fmt.Println("# Top Level Methods")
		for _, method := range topLevelMethods {
			fmt.Println(method)
		}
	}
}

func printLlmNavDetail(target string) {
	for _, sig := range base.GetSortedTSignaturesByClass() {
		if sig.Frame == "Builtin" {
			continue
		}

		if !sigHasCallPoints(sig) {
			continue
		}

		class := sig.Class
		if class == "" {
			class = "Object"
		}

		if class != target && sig.Method != target {
			continue
		}

		if sig.Document == "" {
			sig.Document = "<no document>"
		}

		fmt.Println("## " + sig.GetPrintDetail())
		endRow := findEndRow(sig.Frame, sig.Class, sig.Method)
		fmt.Println("- file: " + sig.FileName + ":" + strconv.Itoa(sig.Row) + "-" + strconv.Itoa(endRow))
		fmt.Println("- document: " + sig.Document)

		points := base.MethodCallPoint[sig.Frame+sig.Class+sig.Method]

		if len(points) == 0 {
			fmt.Println("- callers: none")
		} else {
			fmt.Println("- callers:")
			for _, point := range points {
				if point.CallerMethod != "" {
					fmt.Println("  - method: " + point.CallerMethod)
				} else {
					fmt.Println("  - method: " + "top level")
				}
				if point.CallerClass != "" {
					fmt.Println("    - class: " + point.CallerClass)
				} else {
					fmt.Println("    - class: none")
				}
				fmt.Println("    - call point: " + point.Point)
			}
		}
		fmt.Printf("  - total callers: %d\n", len(points))

		callees := base.MethodCalleePoint[sig.Frame+sig.Class+sig.Method]
		printedCallees := 0

		if len(callees) > 0 {
			fmt.Println("- callees:")
		} else {
			fmt.Println("- callees: none")
		}

		for _, cp := range callees {
			definePoint, ok := findDefinePoint(cp.CalleeFrame, cp.CalleeClass, cp.CalleeMethod)
			if !ok {
				continue
			}
			fmt.Println("  - method: " + cp.CalleeMethod)
			if cp.CalleeClass != "" {
				fmt.Println("    - class: " + cp.CalleeClass)
			} else {
				fmt.Println("    - class: none")
			}
			fmt.Println("    - define point: " + definePoint)
			printedCallees++
		}
		fmt.Printf("  - total callees: %d\n", printedCallees)

		printSpecialComments(sig.FileName, sig.Row, endRow)

		fmt.Println()
	}
}

func findDefinePoint(frame, class, method string) (string, bool) {
	for _, article := range eval.DefineInfoArticles {
		if article.Ctx.GetFrame() == frame &&
			article.Ctx.GetClass() == class &&
			article.Ctx.GetMethod() == method {
			return article.P.FileName + ":" + strconv.Itoa(article.DefineRow) + "-" + strconv.Itoa(article.EndRow), true
		}
	}
	return "", false
}

func findEndRow(frame, class, method string) int {
	for _, article := range eval.DefineInfoArticles {
		if article.Ctx.GetFrame() == frame &&
			article.Ctx.GetClass() == class &&
			article.Ctx.GetMethod() == method {
			return article.EndRow
		}
	}
	return 0
}

func findSpecialComments(fileName string, startRow, endRow int) []base.SpecialCodeComment {
	if endRow == 0 {
		return nil
	}
	var result []base.SpecialCodeComment
	for _, c := range base.SpecialCodeComments {
		if c.FileName == fileName && c.Row >= startRow && c.Row <= endRow {
			result = append(result, c)
		}
	}
	return result
}

func printSpecialComments(fileName string, startRow, endRow int) {
	comments := findSpecialComments(fileName, startRow, endRow)
	if len(comments) == 0 {
		return
	}
	fmt.Println("### special comments")
	for _, c := range comments {
		fmt.Println("#### " + c.FileName + ":" + strconv.Itoa(c.Row))
		fmt.Println("- comment: " + c.Document)
		line := readLineFromFile(c.FileName, c.Row)
		if line != "" {
			fmt.Println("```")
			fmt.Println(line)
			fmt.Println("```")
		}
		fmt.Println()
	}
}

func isCommentInAnyFunction(c base.SpecialCodeComment) bool {
	for _, article := range eval.DefineInfoArticles {
		if article.EndRow == 0 {
			continue
		}
		if article.P.FileName == c.FileName && c.Row >= article.DefineRow && c.Row <= article.EndRow {
			return true
		}
	}
	return false
}

func readLineFromFile(fileName string, row int) string {
	f, err := os.Open(fileName)
	if err != nil {
		return ""
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	current := 0
	for scanner.Scan() {
		current++
		if current == row {
			return scanner.Text()
		}
	}
	return ""
}

func PrintAllErrorsForPlugin(p parser.Parser) {
	for _, err := range p.Errors {
		fmt.Println(err)
	}
}

func PrintAllDefinitionsForLsp(p parser.Parser) {
	printDefinitionTarget(
		p.LspSuggestTargetT.DefinedFrame,
		p.LspSuggestTargetT.DefinedClass,
	)

	printMatchingSignatures(p)
	printInheritanceMap()
}

func PrintSuggestionsForLsp(p parser.Parser) {
	targetT := p.LspSuggestTargetT

	switch targetT.GetType() {
	case base.UNION:
		for _, variant := range targetT.GetVariants() {
			for _, sig := range base.GetSortedTSignatures() {
				if isSuggest(variant, sig) {
					printSuggestion(sig.Method, sig.GetPrintDetail(), sig.Document)
				}
			}
		}
	default:
		isPrinted := false

		for _, sig := range base.GetSortedTSignatures() {
			if isSuggestForKernelOrObjectClass(targetT, sig.Class) {
				printSuggestion(sig.Method, sig.GetPrintDetail(), sig.Document)
				isPrinted = true
				continue
			}

			if isSuggest(targetT, sig) {
				printSuggestion(sig.Method, sig.GetPrintDetail(), sig.Document)
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
		if p.LspSuggestTargetT.IsStatic == sig.IsStatic {
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

func printSuggestion(contents, detail string, document string) {
	fmt.Println(
		prefixSignature + contents + separator + detail + separator + document,
	)
}

func isSuggestForKernelOrObjectClass(targetT base.T, sigClass string) bool {
	if len(targetT.ToString()) == 0 {
		return false
	}

	if unicode.IsUpper(rune(targetT.ToString()[0])) {
		return false
	}

	return slices.Contains([]string{"", "Kernel"}, sigClass)
}

func calculateObjectClassAndIsStatic(targetT base.T) (string, bool) {
	if targetT.IsClassType() {
		return targetT.ToString(), true
	}

	target := targetT.ToString()
	beforeCode := targetT.GetBeforeEvaluateCode()

	var isStaticTarget bool

	switch len(beforeCode) {
	case 0:
		isStaticTarget = unicode.IsUpper(rune(target[0]))
	default:
		isStaticTarget = unicode.IsUpper(rune(beforeCode[0]))
	}

	if targetT.GetType() == base.SELF && !isStaticTarget {
		return targetT.DefinedClass, false
	}

	// 1, '1', 1.1, [], {} and more...
	switch targetT.GetType() {
	case base.INT, base.FLOAT, base.ARRAY, base.HASH, base.STRING, base.OBJECT:
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

	if sig.IsPrivate && sig.Class != targetT.DefinedClass {
		return false
	}

	if sig.Class == targetT.DefinedClass && sig.IsStatic == targetT.IsStatic {
		return true
	}

	if isParentClass(sig, targetT.DefinedFrame, targetT.DefinedClass, targetT.IsStatic, false, false) {
		return true
	}

	if isStaticTarget != sig.IsStatic {
		return false
	}

	if sig.Class == objectClass {
		return true
	}

	return isParentClass(sig, targetT.GetFrame(), objectClass, isStaticTarget, false, false)
}

func isParentClass(
	sig base.Sig,
	frame, class string,
	isStaticTarget bool,
	isExtend bool,
	isInclude bool,
) bool {

	if isExtend && !isStaticTarget {
		return false
	}

	if isInclude && isStaticTarget {
		return false
	}

	if sig.IsStatic != isStaticTarget {
		return false
	}

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
		if isParentClass(sig, parentNode.Frame, parentNode.Class, isStaticTarget, parentNode.IsExtend, parentNode.IsInclude) {
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

func PrintHover(p parser.Parser) {
	targetT := base.GlobT

	for _, sig := range base.GetSortedTSignatures() {
		if targetT.DefinedClass == sig.Class && targetT.GetMethodName() == sig.Method {
			printSuggestion(sig.Method, sig.GetPrintDetail(), sig.Document)
		}
	}
}

func PrintVersion() {
	fmt.Println("ruby-ti version " + Version)
}
