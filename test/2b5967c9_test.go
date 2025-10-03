package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test2b5967c9(t *testing.T) {
	cmd := exec.Command("../ti", "./2b5967c9.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./2b5967c9.rb:::3:::type mismatch: expected Integer, but got String for Integer.to_s"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
