package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test6d28ed59(t *testing.T) {
	cmd := exec.Command("../ti", "./6d28ed59.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./6d28ed59.rb:::9:::type mismatch: expected Block, but got Integer for method_with_block"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
