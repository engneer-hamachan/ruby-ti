package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test7afbfa1b(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./7afbfa1b.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./7afbfa1b.rb:::32:::String"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
