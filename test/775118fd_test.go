package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test775118fd(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./775118fd.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./775118fd.rb:::6:::NilClass"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
