package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestE36c5a34(t *testing.T) {
	cmd := exec.Command("../ti", "./e36c5a34.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./e36c5a34.rb:::12:::type mismatch: expected Union<Integer Float>, but got String for Integer.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
