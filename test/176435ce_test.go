package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test176435ce(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./176435ce.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./176435ce.rb:::6:::Array<Array<String> Integer>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
