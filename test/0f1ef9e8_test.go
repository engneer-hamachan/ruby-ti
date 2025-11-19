package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test0f1ef9e8(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./0f1ef9e8.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./0f1ef9e8.rb:::4:::Bool"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
