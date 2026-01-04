package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestAa39f42b(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./aa39f42b.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := ``

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
