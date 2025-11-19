package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestF4a18c4f(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./f4a18c4f.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./f4a18c4f.rb:::2:::Hash`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
