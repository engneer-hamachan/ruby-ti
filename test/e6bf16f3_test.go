package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestE6bf16f3(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./e6bf16f3.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./e6bf16f3.rb:::26:::String`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
