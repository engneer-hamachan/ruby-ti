package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test6bd1930d(t *testing.T) {
	cmd := exec.Command("../ti", "./6bd1930d.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./6bd1930d.rb:::13:::Integer`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
