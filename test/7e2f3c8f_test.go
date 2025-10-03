package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test7e2f3c8f(t *testing.T) {
	cmd := exec.Command("../ti", "./7e2f3c8f.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./7e2f3c8f.rb:::10:::Hash"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
