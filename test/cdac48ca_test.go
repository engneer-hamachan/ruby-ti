package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestCdac48ca(t *testing.T) {
	cmd := exec.Command("../ti", "./cdac48ca.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./cdac48ca.rb::2::type mismatch: expected String, but got Integer for String.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
