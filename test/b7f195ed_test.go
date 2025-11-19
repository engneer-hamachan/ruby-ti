package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestB7f195ed(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./b7f195ed.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := ""

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
