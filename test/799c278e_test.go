package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test799c278e(t *testing.T) {
	cmd := exec.Command("../ti", "./799c278e.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := ""

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
