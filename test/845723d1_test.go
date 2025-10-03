package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test845723d1(t *testing.T) {
	cmd := exec.Command("../ti", "./845723d1.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./845723d1.rb:::17:::Integer"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
