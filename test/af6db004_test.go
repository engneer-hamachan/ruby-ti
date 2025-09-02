package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestAf6db004(t *testing.T) {
	cmd := exec.Command("../ti", "./af6db004.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./af6db004.rb::9::method '+' is not defined for Identifier"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
