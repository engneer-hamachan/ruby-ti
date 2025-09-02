package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestFfa16881(t *testing.T) {
	cmd := exec.Command("../ti", "./ffa16881.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./ffa16881.rb::9::method '+' is not defined for Identifier"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
