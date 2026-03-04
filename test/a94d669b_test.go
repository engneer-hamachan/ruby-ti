package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestA94d669b(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./a94d669b.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./a94d669b.rb:::4:::Integer`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
