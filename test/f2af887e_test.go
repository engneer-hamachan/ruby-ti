package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestF2af887e(t *testing.T) {
	cmd := exec.Command("../ti", "./f2af887e.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./f2af887e.rb:::4:::type mismatch: expected Union<Integer Float>, but got String for Integer.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
