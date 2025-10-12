package lsp

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"
)

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
