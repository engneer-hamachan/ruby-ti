package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestD5b11546(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./d5b11546.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./d5b11546.rb:::7:::NilClass`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
