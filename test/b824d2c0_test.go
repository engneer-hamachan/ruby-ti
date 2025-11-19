package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestB824d2c0(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./b824d2c0.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := ``

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
