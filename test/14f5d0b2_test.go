package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test14f5d0b2(t *testing.T) {
	cmd := exec.Command("../ti", "./14f5d0b2.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./14f5d0b2.rb:::22:::type mismatch: expected String, but got Integer for String.+\n./14f5d0b2.rb:::23:::method 'test2' is not defined for MatzSanArigato"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
