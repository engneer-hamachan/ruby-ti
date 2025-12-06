package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test75351f34(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./75351f34.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := ``

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
