package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test308971d3(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./308971d3.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./308971d3.rb:::1:::untyped
./308971d3.rb:::3:::untyped`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
