package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test4a814041(t *testing.T) {
	cmd := exec.Command("../ti", "./4a814041.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./4a814041.rb::9::Integer"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
