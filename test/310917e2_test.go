package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test310917e2(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./310917e2.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./310917e2.rb:::5:::untyped`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
