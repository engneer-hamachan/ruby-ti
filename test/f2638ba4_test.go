package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestF2638ba4(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./f2638ba4.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := ""

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
