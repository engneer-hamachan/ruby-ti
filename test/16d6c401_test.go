package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test16d6c401(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./16d6c401.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./16d6c401.rb:::3:::type mismatch: expected String, but got Integer for String.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
