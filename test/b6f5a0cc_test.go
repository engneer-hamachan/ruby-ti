package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestB6f5a0cc(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./b6f5a0cc.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := ``

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
