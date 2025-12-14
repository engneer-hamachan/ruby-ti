package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestCdf98a57(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./cdf98a57.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./cdf98a57.rb:::18:::Union<Integer String>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
