package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestA16beb45(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./a16beb45.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./a16beb45.rb:::17:::[]= is not defined method\n./a16beb45.rb:::19:::Integer"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
