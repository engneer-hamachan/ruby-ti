package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestB5b51ac4(t *testing.T) {
	cmd := exec.Command("../ti", "./b5b51ac4.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./b5b51ac4.rb:::11:::Integer"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
