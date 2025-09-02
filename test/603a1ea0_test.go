package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test603a1ea0(t *testing.T) {
	cmd := exec.Command("../ti", "./603a1ea0.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./603a1ea0.rb::2::expected symbol, but got 'fuga'\n./603a1ea0.rb::17::type mismatch: expected String, but got Integer for String.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
