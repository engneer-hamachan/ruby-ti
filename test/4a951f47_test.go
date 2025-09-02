package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test4a951f47(t *testing.T) {
	cmd := exec.Command("../ti", "./4a951f47.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := ""

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
