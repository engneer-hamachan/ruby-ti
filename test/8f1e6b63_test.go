package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test8f1e6b63(t *testing.T) {
	cmd := exec.Command("../ti", "./8f1e6b63.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./8f1e6b63.rb::4::type mismatch: expected String, but got Integer for String.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
