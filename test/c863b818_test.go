package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestC863b818(t *testing.T) {
	cmd := exec.Command("../ti", "./c863b818.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := ""

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
