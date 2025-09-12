package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test9d0f1c55(t *testing.T) {
	cmd := exec.Command("../ti", "./9d0f1c55.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := ``

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
