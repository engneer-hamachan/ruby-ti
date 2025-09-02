package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestAeeffde5(t *testing.T) {
	cmd := exec.Command("../ti", "./aeeffde5.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./aeeffde5.rb::13::method '+' is not defined for Identifier"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
