package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test166966d1(t *testing.T) {
	cmd := exec.Command("../ti", "./166966d1.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./166966d1.rb::13::method '+' is not defined for Identifier\n./166966d1.rb::17::String"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
