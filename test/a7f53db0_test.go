package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestA7f53db0(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./a7f53db0.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := ""

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
