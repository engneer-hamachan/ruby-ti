package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestDfe9fa87(t *testing.T) {
	cmd := exec.Command("../ti", "./dfe9fa87.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./dfe9fa87.rb:::7:::method '+' is not defined for Nil"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
