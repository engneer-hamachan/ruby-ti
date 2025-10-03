package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestA3fa7d9f(t *testing.T) {
	cmd := exec.Command("../ti", "./a3fa7d9f.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./a3fa7d9f.rb:::18:::Integer`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
