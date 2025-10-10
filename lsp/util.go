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
			base.TSignatures = append(base.TSignatures, base.Sig{
				Contents: methodName,
				Detail:   detail,
			})
		}
	}
}
