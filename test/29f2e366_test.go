package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test29f2e366(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./29f2e366.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./29f2e366.rb:::12:::Integer`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
