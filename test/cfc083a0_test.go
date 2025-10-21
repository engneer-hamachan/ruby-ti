package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestCfc083a0(t *testing.T) {
	cmd := exec.Command("../ti", "./cfc083a0.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./cfc083a0.rb:::4:::Union<Integer String>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
