package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test14c19b0b(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./14c19b0b.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./14c19b0b.rb:::11:::String"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
