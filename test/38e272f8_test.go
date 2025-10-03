package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test38e272f8(t *testing.T) {
	cmd := exec.Command("../ti", "./38e272f8.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./38e272f8.rb:::2:::Integer"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
