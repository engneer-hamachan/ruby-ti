package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test0378be67(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./0378be67.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := ``

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
