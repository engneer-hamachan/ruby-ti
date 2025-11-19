package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestA3997745(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./a3997745.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./a3997745.rb:::15:::Integer`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
