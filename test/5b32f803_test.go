package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test5b32f803(t *testing.T) {
	cmd := exec.Command("../ti", "./5b32f803.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := ``

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
