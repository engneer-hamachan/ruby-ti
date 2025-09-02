package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test7f266531(t *testing.T) {
	cmd := exec.Command("../ti", "./7f266531.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./7f266531.rb::1::type mismatch: expected Float, but got String for Math.sin\n./7f266531.rb::1::type mismatch: expected Integer, but got String for String.*"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
