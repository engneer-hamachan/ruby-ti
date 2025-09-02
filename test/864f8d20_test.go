package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test864f8d20(t *testing.T) {
	cmd := exec.Command("../ti", "./864f8d20.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./864f8d20.rb::6::type mismatch: expected String, but got Bool for test"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
