package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestEa770095(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./ea770095.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./ea770095.rb:::2:::NilClass`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
