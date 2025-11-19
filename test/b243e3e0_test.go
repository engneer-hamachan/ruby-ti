package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestB243e3e0(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./b243e3e0.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./b243e3e0.rb:::2:::Hash`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
