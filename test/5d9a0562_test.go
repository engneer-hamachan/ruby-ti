package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test5d9a0562(t *testing.T) {
	cmd := exec.Command("../ti", "./5d9a0562.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./5d9a0562.rb:::12:::type mismatch: expected String, but got Integer for String.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
