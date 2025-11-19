package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test98701f6c(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./98701f6c.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./98701f6c.rb:::4:::untyped`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
