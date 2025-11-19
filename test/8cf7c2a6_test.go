package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test8cf7c2a6(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./8cf7c2a6.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./8cf7c2a6.rb:::11:::Point`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
