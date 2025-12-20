package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test20eb8f52(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./20eb8f52.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := ``

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
