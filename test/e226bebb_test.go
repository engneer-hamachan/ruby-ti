package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestE226bebb(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./e226bebb.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := ``

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
