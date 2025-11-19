package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestE0b24309(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./e0b24309.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./e0b24309.rb:::5:::Bot"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
