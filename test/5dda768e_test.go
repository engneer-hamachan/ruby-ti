package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test5dda768e(t *testing.T) {
	cmd := exec.Command("../ti", "./5dda768e.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./5dda768e.rb:::5:::type mismatch: expected Union<Integer Float>, but got String for Integer.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
