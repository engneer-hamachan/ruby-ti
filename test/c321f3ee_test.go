package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestC321f3ee(t *testing.T) {
	cmd := exec.Command("../ti", "./c321f3ee.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./c321f3ee.rb:::16:::Integer"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
