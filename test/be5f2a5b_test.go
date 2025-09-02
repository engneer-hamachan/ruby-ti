package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestBe5f2a5b(t *testing.T) {
	cmd := exec.Command("../ti", "./be5f2a5b.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := ""

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
