package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test14d57ed3(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./14d57ed3.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./14d57ed3.rb:::7:::type mismatch: expected String, but got Integer for String.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
