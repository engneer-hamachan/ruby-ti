package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestEe885b00(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./ee885b00.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./ee885b00.rb:::3:::Integer"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
