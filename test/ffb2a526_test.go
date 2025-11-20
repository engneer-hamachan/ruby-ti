package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestFfb2a526(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./ffb2a526.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./ffb2a526.rb:::3:::String
./ffb2a526.rb:::8:::Integer`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
