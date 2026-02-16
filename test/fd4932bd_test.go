package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestFd4932bd(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./fd4932bd.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./fd4932bd.rb:::3:::type mismatch: expected String, but got Integer for String.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
