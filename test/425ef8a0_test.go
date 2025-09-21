package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test425ef8a0(t *testing.T) {
	cmd := exec.Command("../ti", "./425ef8a0.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./425ef8a0.rb::20::String`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
