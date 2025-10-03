package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestD12defc0(t *testing.T) {
	cmd := exec.Command("../ti", "./d12defc0.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./d12defc0.rb:::10:::Union<Nil Integer>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
