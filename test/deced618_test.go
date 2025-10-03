package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestDeced618(t *testing.T) {
	cmd := exec.Command("../ti", "./deced618.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./deced618.rb:::17:::Integer`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
