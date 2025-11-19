package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestEd339c31(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./ed339c31.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := ``

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
