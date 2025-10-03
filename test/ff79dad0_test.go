package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestFf79dad0(t *testing.T) {
	cmd := exec.Command("../ti", "./ff79dad0.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./ff79dad0.rb:::11:::Union<Bot Integer>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
