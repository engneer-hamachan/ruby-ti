package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestC67708a5(t *testing.T) {
	cmd := exec.Command("../ti", "./c67708a5.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./c67708a5.rb::9::type mismatch: expected Union<Integer Float>, but got String for Integer.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
