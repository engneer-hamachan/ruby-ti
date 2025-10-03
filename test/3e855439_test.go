package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test3e855439(t *testing.T) {
	cmd := exec.Command("../ti", "./3e855439.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./3e855439.rb:::23:::type mismatch: expected Union<Integer Float>, but got String for Integer.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
