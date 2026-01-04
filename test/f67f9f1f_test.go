package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestF67f9f1f(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./f67f9f1f.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./f67f9f1f.rb:::7:::NilClass`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
