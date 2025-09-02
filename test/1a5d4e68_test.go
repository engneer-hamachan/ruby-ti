package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test1a5d4e68(t *testing.T) {
	cmd := exec.Command("../ti", "./1a5d4e68.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./1a5d4e68.rb::8::Integer"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
