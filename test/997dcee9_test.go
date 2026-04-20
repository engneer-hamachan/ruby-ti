package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test997dcee9(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./997dcee9.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./997dcee9.rb:::8:::Integer`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
