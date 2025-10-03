package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test298174f8(t *testing.T) {
	cmd := exec.Command("../ti", "./298174f8.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./298174f8.rb:::16:::method 'test' is not defined for Fuga`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
