package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test87ed0d11(t *testing.T) {
	cmd := exec.Command("../ti", "./87ed0d11.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./87ed0d11.rb::4::Nil`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
