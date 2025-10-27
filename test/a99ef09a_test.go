package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestA99ef09a(t *testing.T) {
	cmd := exec.Command("../ti", "./a99ef09a.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./a99ef09a.rb:::5:::Integer`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
