package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test15fe3c1c(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./15fe3c1c.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./15fe3c1c.rb:::9:::Integer`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
