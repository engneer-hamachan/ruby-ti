package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test78c0e488(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./78c0e488.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./78c0e488.rb:::11:::Union<Nil String>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
