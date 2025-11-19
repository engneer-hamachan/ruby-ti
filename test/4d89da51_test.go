package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test4d89da51(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./4d89da51.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./4d89da51.rb:::3:::Integer`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
