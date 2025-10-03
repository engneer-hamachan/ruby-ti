package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test63b1e74a(t *testing.T) {
	cmd := exec.Command("../ti", "./63b1e74a.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./63b1e74a.rb:::3:::too many arguments for Integer.to_s"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
