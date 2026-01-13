package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestA31d4e79(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./a31d4e79.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./a31d4e79.rb:::5:::untyped
./a31d4e79.rb:::6:::untyped`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
