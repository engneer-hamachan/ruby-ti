package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestBf2e01bc(t *testing.T) {
	cmd := exec.Command("../ti", "./bf2e01bc.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./bf2e01bc.rb::6::type mismatch: expected Union<Integer Float>, but got String for Float.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
