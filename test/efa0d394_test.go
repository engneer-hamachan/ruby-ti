package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestEfa0d394(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./efa0d394.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./efa0d394.rb:::3:::type mismatch: expected String, but got Integer for String.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
