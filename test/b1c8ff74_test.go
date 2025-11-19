package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestB1c8ff74(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./b1c8ff74.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := ``

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
