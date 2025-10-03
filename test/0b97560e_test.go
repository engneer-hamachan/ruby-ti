package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test0b97560e(t *testing.T) {
	cmd := exec.Command("../ti", "./0b97560e.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./0b97560e.rb:::11:::Integer"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
