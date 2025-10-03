package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestE206d98a(t *testing.T) {
	cmd := exec.Command("../ti", "./e206d98a.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./e206d98a.rb:::1:::type mismatch: expected Union<Integer Float>, but got String for Integer.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
