package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestFb1d5905(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./fb1d5905.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./fb1d5905.rb:::16:::type mismatch: expected Union<Integer Float>, but got String for Integer.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
