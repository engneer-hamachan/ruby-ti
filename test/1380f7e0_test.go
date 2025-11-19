package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test1380f7e0(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./1380f7e0.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./1380f7e0.rb:::5:::type mismatch: expected Block, but got String for Array.each"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
