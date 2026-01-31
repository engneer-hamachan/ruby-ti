package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestAa56faee(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./aa56faee.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := ``

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
