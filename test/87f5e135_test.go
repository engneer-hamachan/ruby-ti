package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test87f5e135(t *testing.T) {
	cmd := exec.Command("../ti", "./87f5e135.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./87f5e135.rb::17::String`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
