package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestF1446133(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./f1446133.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := ``

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
