package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test26f0216a(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./26f0216a.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := ``

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
