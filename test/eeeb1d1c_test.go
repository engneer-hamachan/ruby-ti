package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestEeeb1d1c(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./eeeb1d1c.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := ``

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
