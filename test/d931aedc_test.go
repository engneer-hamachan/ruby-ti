package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestD931aedc(t *testing.T) {
	cmd := exec.Command("../ti", "./d931aedc.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := ""

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
