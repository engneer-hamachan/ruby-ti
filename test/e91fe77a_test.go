package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestE91fe77a(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./e91fe77a.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./e91fe77a.rb:::13:::Integer"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
