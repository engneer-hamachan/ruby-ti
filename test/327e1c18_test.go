package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test327e1c18(t *testing.T) {
	cmd := exec.Command("../ti", "./327e1c18.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./327e1c18.rb::11::type mismatch: expected Union<Integer Float>, but got String for Integer.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
