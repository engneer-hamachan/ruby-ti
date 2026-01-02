package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestE95a39a8(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./e95a39a8.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./e95a39a8.rb:::11:::NilClass"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
