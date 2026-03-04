package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestBcf72b3b(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./bcf72b3b.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./bcf72b3b.rb:::11:::Union<Float Integer>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
