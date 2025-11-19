package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestB06a5d8e(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./b06a5d8e.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./b06a5d8e.rb:::16:::type mismatch: expected String, but got Integer for String.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
