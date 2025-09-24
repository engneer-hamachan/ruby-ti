package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test94b9f8b0(t *testing.T) {
	cmd := exec.Command("../ti", "./94b9f8b0.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./94b9f8b0.rb::31::String`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
