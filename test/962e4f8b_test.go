package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test962e4f8b(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./962e4f8b.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./962e4f8b.rb:::7:::c: is not defined expected (Integer, a: Integer, b: Integer, c: unknown)\n./962e4f8b.rb:::7:::too few arguments for test expected (Integer, a: Integer, b: Integer, c: unknown)"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
