package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestD1856ea4(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./d1856ea4.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./d1856ea4.rb:::5:::String`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
