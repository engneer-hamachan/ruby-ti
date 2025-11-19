package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test86a4b4d1(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./86a4b4d1.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./86a4b4d1.rb:::128:::Hash"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
