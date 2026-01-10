package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestC5a95372(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./c5a95372.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := ``

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
