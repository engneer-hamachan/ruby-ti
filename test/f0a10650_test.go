package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestF0a10650(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./f0a10650.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./f0a10650.rb:::4:::Nil"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
