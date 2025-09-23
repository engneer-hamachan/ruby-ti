package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestC0c07166(t *testing.T) {
	cmd := exec.Command("../ti", "./c0c07166.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./c0c07166.rb::1::Symbol`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
