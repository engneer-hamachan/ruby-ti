package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestC4fec57c(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./c4fec57c.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./c4fec57c.rb:::7:::Array<Integer String>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
