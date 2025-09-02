package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test55268bb9(t *testing.T) {
	cmd := exec.Command("../ti", "./55268bb9.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./55268bb9.rb::1::Float"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
