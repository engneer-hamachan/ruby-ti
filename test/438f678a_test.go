package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test438f678a(t *testing.T) {
	cmd := exec.Command("../ti", "./438f678a.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./438f678a.rb:::5:::untyped"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
