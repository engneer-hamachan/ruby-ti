package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test35f7c829(t *testing.T) {
	cmd := exec.Command("../ti", "./35f7c829.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./35f7c829.rb:::15:::method 'test' is not defined for Fuga`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
