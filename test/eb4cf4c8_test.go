package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestEb4cf4c8(t *testing.T) {
	cmd := exec.Command("../ti", "./eb4cf4c8.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./eb4cf4c8.rb::4::method '+' is not defined for Identifier\n./eb4cf4c8.rb::7::too few arguments for test expected (Integer, *untyped, ?)"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
