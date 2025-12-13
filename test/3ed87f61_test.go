package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test3ed87f61(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./3ed87f61.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./3ed87f61.rb:::19:::Integer`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
