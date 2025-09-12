package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test7a5b4733(t *testing.T) {
	cmd := exec.Command("../ti", "./7a5b4733.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := ``

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
