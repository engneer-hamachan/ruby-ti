package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestD63e6cb4(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./d63e6cb4.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./d63e6cb4.rb:::37:::type mismatch: expected String, but got Integer for String.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
