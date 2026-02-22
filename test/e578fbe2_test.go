package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestE578fbe2(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./e578fbe2.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./e578fbe2.rb:::4:::Integer`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
