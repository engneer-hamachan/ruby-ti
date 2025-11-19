package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestAb3c6d69(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./ab3c6d69.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./ab3c6d69.rb:::9:::String`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
