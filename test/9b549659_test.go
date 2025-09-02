package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test9b549659(t *testing.T) {
	cmd := exec.Command("../ti", "./9b549659.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./9b549659.rb::14::String"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
