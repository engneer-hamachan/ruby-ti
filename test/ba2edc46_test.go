package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestBa2edc46(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./ba2edc46.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./ba2edc46.rb:::2:::Integer`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
