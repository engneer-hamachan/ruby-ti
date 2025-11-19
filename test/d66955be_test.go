package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestD66955be(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./d66955be.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := ""

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
