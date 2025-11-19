package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test27390773(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./27390773.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./27390773.rb:::5:::type mismatch: expected Union<Integer Float>, but got String for Integer.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
