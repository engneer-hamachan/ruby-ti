package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test8904c9bd(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./8904c9bd.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./8904c9bd.rb:::6:::Integer"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
