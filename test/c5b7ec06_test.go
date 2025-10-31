package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestC5b7ec06(t *testing.T) {
	cmd := exec.Command("../ti", "./c5b7ec06.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./c5b7ec06.rb:::7:::instance method '+' is not defined for Nil"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
