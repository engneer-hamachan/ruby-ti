package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test75c8afb8(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./75c8afb8.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./75c8afb8.rb:::9:::Union<String NilClass>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
