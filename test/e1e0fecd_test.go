package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestE1e0fecd(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./e1e0fecd.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./e1e0fecd.rb:::7:::Hash`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
