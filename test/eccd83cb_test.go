package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestEccd83cb(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./eccd83cb.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./eccd83cb.rb:::7:::type mismatch: expected String, but got Integer for String.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
