package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test311d05cb(t *testing.T) {
	cmd := exec.Command("../ti", "./311d05cb.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./311d05cb.rb::2::type mismatch: expected Union<Integer Float>, but got String for Integer.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
