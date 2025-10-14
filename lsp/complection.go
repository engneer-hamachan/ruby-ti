package lsp

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"ti/base"
	"time"
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

func analyzeContent(content string, line uint32) error {
	responseSignatures = nil

	content = removeTaiilDot(content, line)

	tmpFile, err := os.CreateTemp("", "ruby-ti-lsp-*.rb")
	if err != nil {
		return nil
	}

	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	if _, err := tmpFile.WriteString(content); err != nil {
		return nil
	}

	tmpFile.Close()

	ctx, cancel :=
		context.WithTimeout(context.Background(), 1000*time.Millisecond)

	defer cancel()

	cmd :=
		exec.CommandContext(
			ctx,
			"ti",
			tmpFile.Name(),
			"-a",
			fmt.Sprintf("%d", line+1),
		)

	output, err := cmd.Output()
	if err != nil {
		return nil
	}

	setTSignatures(output)

	return nil
}
