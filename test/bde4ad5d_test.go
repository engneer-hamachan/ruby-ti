package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestBde4ad5d(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./bde4ad5d.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./bde4ad5d.rb:::11:::untyped"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
