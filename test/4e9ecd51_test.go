package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test4e9ecd51(t *testing.T) {
	cmd := exec.Command("../ti", "./4e9ecd51.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := ``

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
