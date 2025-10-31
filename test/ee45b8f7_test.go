package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestEe45b8f7(t *testing.T) {
	cmd := exec.Command("../ti", "./ee45b8f7.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./ee45b8f7.rb:::9:::instance method '+' is not defined for Nil"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
