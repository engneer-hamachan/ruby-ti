package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestAa239304(t *testing.T) {
	cmd := exec.Command("../ti", "./aa239304.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := ""

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
