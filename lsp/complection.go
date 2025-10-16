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

func getSignatures(cmdOutput []byte) []base.Sig {
	methodSet := make(map[string]bool)
	scanner := bufio.NewScanner(strings.NewReader(string(cmdOutput)))

	var responseSignatures []base.Sig

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
				Method: methodName,
				Detail: detail,
			})
		}
	}

	return responseSignatures
}

func findComplection(content string, line uint32) []base.Sig {
	content = removeTaiilDot(content, line)

	tmpFile, err := os.CreateTemp("", "ruby-ti-lsp-*.rb")
	if err != nil {
		return []base.Sig{}
	}

	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	if _, err := tmpFile.WriteString(content); err != nil {
		return []base.Sig{}
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
		return []base.Sig{}
	}

	return getSignatures(output)
}
