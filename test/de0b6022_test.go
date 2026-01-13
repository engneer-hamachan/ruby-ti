package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestDe0b6022(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./de0b6022.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := ``

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
