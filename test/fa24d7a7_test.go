package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestFa24d7a7(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./fa24d7a7.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := ``

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
