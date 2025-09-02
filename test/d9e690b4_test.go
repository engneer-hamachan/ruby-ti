package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestD9e690b4(t *testing.T) {
	cmd := exec.Command("../ti", "./d9e690b4.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := ""

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
