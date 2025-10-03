package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test3c1e8e6a(t *testing.T) {
	cmd := exec.Command("../ti", "./3c1e8e6a.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./3c1e8e6a.rb:::9:::type mismatch: expected Union<Integer Float>, but got String for Integer.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
