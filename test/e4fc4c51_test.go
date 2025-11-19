package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestE4fc4c51(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./e4fc4c51.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./e4fc4c51.rb:::9:::Union<Integer Nil>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
