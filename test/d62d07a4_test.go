package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestD62d07a4(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./d62d07a4.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./d62d07a4.rb:::6:::Array<String Symbol Integer>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
