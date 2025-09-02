package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestDa6c0758(t *testing.T) {
	cmd := exec.Command("../ti", "./da6c0758.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := ""

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
