package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestD5ecf019(t *testing.T) {
	cmd := exec.Command("../ti", "./d5ecf019.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./d5ecf019.rb::3::type mismatch: expected Union<Integer Float>, but got String for Integer.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
