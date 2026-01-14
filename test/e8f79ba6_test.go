package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestE8f79ba6(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./e8f79ba6.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./e8f79ba6.rb:::10:::untyped`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
