package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestB0259e64(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./b0259e64.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./b0259e64.rb:::17:::Integer`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
