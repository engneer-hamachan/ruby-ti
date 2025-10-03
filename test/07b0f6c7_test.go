package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test07b0f6c7(t *testing.T) {
	cmd := exec.Command("../ti", "./07b0f6c7.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./07b0f6c7.rb:::1:::Integer"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
