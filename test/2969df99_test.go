package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test2969df99(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./2969df99.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./2969df99.rb:::26:::Integer"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
