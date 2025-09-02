package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestBdcd1bce(t *testing.T) {
	cmd := exec.Command("../ti", "./bdcd1bce.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./bdcd1bce.rb::6::type mismatch: expected Union<Integer Float>, but got Bool for Integer.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
