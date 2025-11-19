package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test4dcc7d3e(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./4dcc7d3e.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./4dcc7d3e.rb:::3:::Union<Integer Array<untyped>>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
