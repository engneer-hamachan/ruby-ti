package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test18db56e0(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./18db56e0.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./18db56e0.rb:::31:::type mismatch: expected String, but got Integer for String.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
