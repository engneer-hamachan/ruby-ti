package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test41df5ce0(t *testing.T) {
	cmd := exec.Command("../ti", "./41df5ce0.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./41df5ce0.rb:::12:::Unknown"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
