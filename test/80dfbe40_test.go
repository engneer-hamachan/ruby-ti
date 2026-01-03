package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test80dfbe40(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./80dfbe40.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := ``

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
