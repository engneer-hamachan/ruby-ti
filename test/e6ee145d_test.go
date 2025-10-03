package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestE6ee145d(t *testing.T) {
	cmd := exec.Command("../ti", "./e6ee145d.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./e6ee145d.rb:::9:::type mismatch: expected Union<Integer Float>, but got String for Integer.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
