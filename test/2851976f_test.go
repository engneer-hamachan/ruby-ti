package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test2851976f(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./2851976f.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := ``

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
