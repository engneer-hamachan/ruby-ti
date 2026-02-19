package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test657dd04b(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./657dd04b.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./657dd04b.rb:::17:::Integer`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
