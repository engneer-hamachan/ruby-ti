package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test9477eea4(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./9477eea4.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./9477eea4.rb:::13:::Integer`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
