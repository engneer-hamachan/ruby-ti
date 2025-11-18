package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test236139b8(t *testing.T) {
	cmd := exec.Command("../ti", "./236139b8.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./236139b8.rb:::8:::User`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
