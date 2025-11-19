package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestA8815041(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./a8815041.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./a8815041.rb:::14:::Symbol"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
