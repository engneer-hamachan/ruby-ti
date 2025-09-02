package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestFf913141(t *testing.T) {
	cmd := exec.Command("../ti", "./ff913141.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./ff913141.rb::2::type mismatch: expected Union<Integer Float>, but got String for Integer.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
