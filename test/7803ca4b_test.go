package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test7803ca4b(t *testing.T) {
	cmd := exec.Command("../ti", "./7803ca4b.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := ""

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
