package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestBeff9b9b(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./beff9b9b.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./beff9b9b.rb:::2:::Integer`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
