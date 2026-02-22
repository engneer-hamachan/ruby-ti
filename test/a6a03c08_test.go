package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestA6a03c08(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./a6a03c08.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./a6a03c08.rb:::2:::Integer`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
