package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestB772d987(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./b772d987.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./b772d987.rb:::11:::type mismatch: expected Union<Integer Float>, but got String for Integer.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
