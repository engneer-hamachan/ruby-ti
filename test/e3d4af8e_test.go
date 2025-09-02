package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestE3d4af8e(t *testing.T) {
	cmd := exec.Command("../ti", "./e3d4af8e.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./e3d4af8e.rb::1::Range"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
