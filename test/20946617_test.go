package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test20946617(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./20946617.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./20946617.rb:::2:::type mismatch: expected Integer, but got String for Test.test"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
