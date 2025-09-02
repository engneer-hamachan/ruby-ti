package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test504a32c0(t *testing.T) {
	cmd := exec.Command("../ti", "./504a32c0.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./504a32c0.rb::16::Integer"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
