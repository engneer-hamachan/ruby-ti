package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestB268e480(t *testing.T) {
	cmd := exec.Command("../ti", "./b268e480.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./b268e480.rb::15::type mismatch: expected Union<Integer Float>, but got String for Integer.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
