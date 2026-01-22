package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test0ebb4ae9(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./0ebb4ae9.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./0ebb4ae9.rb:::12:::Integer`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
