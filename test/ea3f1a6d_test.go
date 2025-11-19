package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestEa3f1a6d(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./ea3f1a6d.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./ea3f1a6d.rb:::6:::String
./ea3f1a6d.rb:::8:::Integer`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
