package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test63185e90(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./63185e90.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./63185e90.rb:::5:::String"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
