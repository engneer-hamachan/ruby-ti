package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test2d05491a(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./2d05491a.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := ``

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
