package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test9e2a206f(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./9e2a206f.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./9e2a206f.rb:::5:::untyped`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
