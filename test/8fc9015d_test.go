package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test8fc9015d(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./8fc9015d.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := ``

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
