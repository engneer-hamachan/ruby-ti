package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestBbd2f302(t *testing.T) {
	cmd := exec.Command("../ti", "./bbd2f302.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./bbd2f302.rb::15::Integer"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
