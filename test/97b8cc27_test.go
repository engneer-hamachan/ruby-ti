package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test97b8cc27(t *testing.T) {
	cmd := exec.Command("../ti", "./97b8cc27.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./97b8cc27.rb:::19:::type mismatch: expected Union<Integer Float>, but got String for Integer.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
