package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestEa9cf3d3(t *testing.T) {
	cmd := exec.Command("../ti", "./ea9cf3d3.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./ea9cf3d3.rb::14::Integer"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
