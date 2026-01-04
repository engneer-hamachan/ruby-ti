package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test187d5798(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./187d5798.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./187d5798.rb:::7:::untyped`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
