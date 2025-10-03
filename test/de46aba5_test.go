package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestDe46aba5(t *testing.T) {
	cmd := exec.Command("../ti", "./de46aba5.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./de46aba5.rb:::7:::Integer"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
