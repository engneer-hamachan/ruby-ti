package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestAe2bb379(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./ae2bb379.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./ae2bb379.rb:::11:::String"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
