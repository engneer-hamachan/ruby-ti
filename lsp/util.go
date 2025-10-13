package lsp

import (
	"bufio"
	"strings"
	"ti/base"
)

func removeTaiilDot(content string, line uint32) string {
	lines := strings.Split(content, "\n")
	if int(line) < len(lines) {
		currentLine := lines[line]
		trimmed := strings.TrimSpace(currentLine)
		if strings.HasSuffix(trimmed, ".") {
			lines[line] = strings.TrimSuffix(currentLine, ".")
			content = strings.Join(lines, "\n")
		}
	}

	return content
}

func setTSignatures(cmdOutput []byte) {
	methodSet := make(map[string]bool)
	scanner := bufio.NewScanner(strings.NewReader(string(cmdOutput)))
	for scanner.Scan() {
		line := scanner.Text()
		line, ok := strings.CutPrefix(line, "%")
		if !ok {
			continue
		}

		parts := strings.SplitN(line, ":::", 2)
		if len(parts) < 2 {
			continue
		}
		methodName := parts[0]
		detail := parts[1]

		if !methodSet[detail] {
			methodSet[detail] = true
			responseSignatures = append(responseSignatures, base.Sig{
				Contents: methodName,
				Detail:   detail,
			})
		}
	}
}

// extractMethodName extracts method name from line at cursor position
// Examples: "h.test 1" -> "test", "test 1" -> "test"
func extractMethodName(line string, col int) string {
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

	if start == end {
		return ""
	}

	return line[start:end]
}

// extractTargetForPrefix extracts target for -a option
// Examples: "h.test 1" -> "h", "test 1" -> "test"
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
		// h.test の場合、h を返す
		// ドットより前の単語を探す
		dotStart := dotPos
		for dotStart > 0 && isWordChar(line[dotStart-1]) {
			dotStart--
		}
		if dotStart < dotPos {
			return line[dotStart:dotPos]
		}
	}

	// ドットがない場合は、メソッド名だけを返す
	return line[start:end]
}

// isWordChar checks if a byte is part of a word (letter, digit, underscore)
func isWordChar(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || (b >= '0' && b <= '9') || b == '_'
}
