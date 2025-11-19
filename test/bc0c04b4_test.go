package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestBc0c04b4(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./bc0c04b4.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := ""

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
