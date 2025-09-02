package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test127eb360(t *testing.T) {
	cmd := exec.Command("../ti", "./127eb360.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./127eb360.rb::3::type mismatch. x is not Array or Hash"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
