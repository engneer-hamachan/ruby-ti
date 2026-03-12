package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestA6729b3f(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./a6729b3f.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./a6729b3f.rb:::9:::Bool`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
