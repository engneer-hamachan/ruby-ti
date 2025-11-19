package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test5cf6cf48(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./5cf6cf48.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./5cf6cf48.rb:::3:::type mismatch: expected Array<untyped>, but got Integer for Array.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
