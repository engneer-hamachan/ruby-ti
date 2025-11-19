package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestA1a28711(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./a1a28711.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./a1a28711.rb:::6:::instance method 'abs' is not defined for String"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
