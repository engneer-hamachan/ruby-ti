package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestEa8e5997(t *testing.T) {
	cmd := exec.Command("../ti", "./ea8e5997.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./ea8e5997.rb:::15:::type mismatch: expected String, but got Integer for String.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
