package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestDc979e82(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./dc979e82.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := ``

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
