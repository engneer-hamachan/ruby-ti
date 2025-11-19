package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test38fc6681(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./38fc6681.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := ``

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
