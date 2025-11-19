package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestBa6d26d3(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./ba6d26d3.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := ""

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
