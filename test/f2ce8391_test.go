package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestF2ce8391(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./f2ce8391.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./f2ce8391.rb:::6:::Integer"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
