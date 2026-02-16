package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestF14eb446(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./f14eb446.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./f14eb446.rb:::8:::type mismatch: expected Union<Integer Float>, but got String for Integer.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
