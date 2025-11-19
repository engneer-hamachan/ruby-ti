package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestF14e7097(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./f14e7097.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./f14e7097.rb:::17:::Integer`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
