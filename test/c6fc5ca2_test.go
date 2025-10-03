package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestC6fc5ca2(t *testing.T) {
	cmd := exec.Command("../ti", "./c6fc5ca2.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./c6fc5ca2.rb:::9:::type mismatch: expected Union<Integer Float>, but got String for Integer.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
