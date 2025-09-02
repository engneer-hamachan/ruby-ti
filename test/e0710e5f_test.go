package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestE0710e5f(t *testing.T) {
	cmd := exec.Command("../ti", "./e0710e5f.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := ""

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
