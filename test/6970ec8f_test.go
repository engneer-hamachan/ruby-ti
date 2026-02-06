package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test6970ec8f(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./6970ec8f.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := ``

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
