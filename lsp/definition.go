package lsp

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strings"
	"ti/base"
	"time"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

// h.test 1 -> test, test 1 -> test, test? -> test?, attr= -> attr=
func extractMethodName(line string, col int) string {
	if col > len(line) {
		col = len(line)
	}

	start := col
	for start > 0 && isWordChar(line[start-1]) {
		start--
	}

	end := col
	for end < len(line) && isWordChar(line[end]) {
		end++
	}

	specialChars := []byte{'?', '!', '='}
	if len(line) > end && slices.Contains(specialChars, line[end]) {
		end++
	}

	if start == end {
		return ""
	}

	return line[start:end]
}

// extractTargetForPrefix extracts target for -a option
// Examples: "h.test 1" -> "h.test", "test 1" -> "test", "h.nil?" -> "h.nil?"
func extractTargetForPrefix(line string, col int) string {
	if col > len(line) {
		col = len(line)
	}

	// カーソル位置の単語の開始位置を探す
	start := col
	for start > 0 && isWordChar(line[start-1]) {
		start--
	}

	// カーソル位置の単語の終了位置を探す
	end := col
	for end < len(line) && isWordChar(line[end]) {
		end++
	}

	// Ruby メソッド名の末尾に ?, !, = がある場合は含める
	if end < len(line) && (line[end] == '?' || line[end] == '!' || line[end] == '=') {
		end++
	}

	if start == end {
		return ""
	}

	// ドットがあるかチェック（start より前を見る）
	dotPos := -1
	for i := start - 1; i >= 0; i-- {
		if line[i] == '.' {
			dotPos = i
			break
		} else if line[i] != ' ' && line[i] != '\t' {
			// ドット以外の文字があったら終了
			break
		}
	}

	if dotPos >= 0 {
		// h.test の場合、h.test 全体を返す
		// ドットより前の単語を探す
		dotStart := dotPos
		for dotStart > 0 && isWordChar(line[dotStart-1]) {
			dotStart--
		}
		if dotStart < dotPos {
			return line[dotStart:end]
		}
	}

	// ドットがない場合は、メソッド名だけを返す
	return line[start:end]
}

// isWordChar checks if a byte is part of a word (letter, digit, underscore)
func isWordChar(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || (b >= '0' && b <= '9') || b == '_'
}

func findDefinition(content string, params *protocol.DefinitionParams) (any, error) {
	codeLines := strings.Split(content, "\n")
	if int(params.Position.Line) >= len(codeLines) {
		return nil, nil
	}

	currentLine := codeLines[params.Position.Line]
	methodName := extractMethodName(currentLine, int(params.Position.Character))
	if methodName == "" {
		return nil, nil
	}

	targetForPrefix :=
		extractTargetForPrefix(currentLine, int(params.Position.Character))

	// ドットが含まれているかチェック（レシーバがあるか）
	hasDot := strings.Contains(targetForPrefix, ".")

	codeLines[params.Position.Line] = targetForPrefix
	modifiedContent := strings.Join(codeLines, "\n")

	// 一時ファイルを作成
	tmpFile, err := os.CreateTemp("", "ruby-ti-lsp-*.rb")
	if err != nil {
		return nil, nil
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	if _, err := tmpFile.WriteString(modifiedContent); err != nil {
		return nil, nil
	}
	tmpFile.Close()

	// ti {file} -a {row} で @prefix を取得
	prefixInfo := getPrefixInfo(tmpFile.Name(), int(params.Position.Line)+1)
	if prefixInfo == "" {
		return nil, nil
	}

	// @{frame}:::{class} からクラス情報を抽出
	parts := strings.SplitN(strings.TrimPrefix(prefixInfo, "@"), ":::", 2)
	if len(parts) < 2 {
		return nil, nil
	}
	frame := parts[0]
	class := parts[1]

	// ti {file} --define で全メソッド定義と継承情報を取得
	definitions, inheritanceMap := getMethodDefinitionsAndInheritance(tmpFile.Name())

	// ドットが入っていなくて、frameがunknownでclassも空だったらtoplevelメソッド
	searchFrame := frame
	searchClass := class
	if !hasDot && frame == "unknown" && class == "" {
		searchFrame = "unknown"
		searchClass = "unknown"
	}

	// マッチするメソッド定義を検索
	// フォーマット: %{frame}:::{class}:::{method}:::{filename}:::{row}
	for _, def := range definitions {
		defParts := strings.SplitN(def, ":::", 5)
		if len(defParts) < 5 {
			continue
		}

		defFrame := defParts[0]
		defClass := defParts[1]
		defMethod := defParts[2]
		defFilename := defParts[3]
		defRow := defParts[4]

		// Check if method matches (including parent class methods)
		if isMethodMatch(defFrame, defClass, searchFrame, searchClass, methodName, defMethod, inheritanceMap) {
			// 定義位置を返す
			var row uint32
			fmt.Sscanf(defRow, "%d", &row)
			if row > 0 {
				row-- // 0-based indexing
			}

			// tmpファイルのパスの場合は、元のファイルのURIを使う
			targetURI := params.TextDocument.URI
			if !strings.Contains(defFilename, "ruby-ti-lsp-") {
				// 別ファイルの定義の場合はそのパスを使う
				targetURI = protocol.DocumentUri("file://" + defFilename)
			}

			location := protocol.Location{
				URI: targetURI,
				Range: protocol.Range{
					Start: protocol.Position{
						Line:      row,
						Character: 0,
					},
					End: protocol.Position{
						Line:      row,
						Character: 0,
					},
				},
			}

			return location, nil
		}
	}

	return nil, nil
}

// getPrefixInfo gets @prefix info using ti -a
func getPrefixInfo(filename string, row int) string {
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()

	cmd := exec.CommandContext(ctx, "ti", filename, "-a", fmt.Sprintf("%d", row))
	output, err := cmd.Output()
	if err != nil {
		return ""
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "@") {
			return line
		}
	}

	return ""
}

// getMethodDefinitionsAndInheritance gets all method definitions and inheritance info using ti --define
func getMethodDefinitionsAndInheritance(filename string) ([]string, map[base.ClassNode][]base.ClassNode) {
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()

	cmd := exec.CommandContext(ctx, "ti", filename, "--define")
	output, err := cmd.Output()
	if err != nil {
		return nil, make(map[base.ClassNode][]base.ClassNode)
	}

	var definitions []string
	inheritanceMap := make(map[base.ClassNode][]base.ClassNode)

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "%") {
			// Method definition
			definitions = append(definitions, strings.TrimPrefix(line, "%"))
		} else if strings.HasPrefix(line, "$") {
			// Inheritance information
			line = strings.TrimPrefix(line, "$")
			parts := strings.SplitN(line, ":::", 4)
			if len(parts) < 4 {
				continue
			}

			childNode := base.ClassNode{Frame: parts[0], Class: parts[1]}
			parentNode := base.ClassNode{Frame: parts[2], Class: parts[3]}

			inheritanceMap[childNode] = append(inheritanceMap[childNode], parentNode)
		}
	}

	return definitions, inheritanceMap
}

// normalizeFrame normalizes empty string and "unknown" to be the same
func normalizeFrame(frame string) string {
	if frame == "" || frame == "unknown" {
		return ""
	}
	return frame
}

// isParentClass checks if parentClass is a parent of childClass
func isParentClass(frame, childClass, parentClass string, inheritanceMap map[base.ClassNode][]base.ClassNode) bool {
	// Try with normalized frame (empty string and "unknown" are treated the same)
	normalizedFrame := normalizeFrame(frame)

	// Try both the original frame and normalized frame
	framesToTry := []string{frame, normalizedFrame}
	if frame == "unknown" {
		framesToTry = append(framesToTry, "")
	} else if frame == "" {
		framesToTry = append(framesToTry, "unknown")
	}

	for _, tryFrame := range framesToTry {
		classNode := base.ClassNode{Frame: tryFrame, Class: childClass}

		for _, parentNode := range inheritanceMap[classNode] {
			if parentNode.Class == parentClass {
				return true
			}

			// Recursively check parent's parents
			if isParentClass(parentNode.Frame, parentNode.Class, parentClass, inheritanceMap) {
				return true
			}
		}
	}

	return false
}

// isMethodMatch checks if the method definition matches the search criteria
// considering inheritance
func isMethodMatch(defFrame, defClass, searchFrame, searchClass, methodName, defMethod string, inheritanceMap map[base.ClassNode][]base.ClassNode) bool {
	if defMethod != methodName {
		return false
	}

	// Exact match
	if defFrame == searchFrame && defClass == searchClass {
		return true
	}

	// Check if defClass is a parent class of searchClass
	return isParentClass(searchFrame, searchClass, defClass, inheritanceMap)
}
