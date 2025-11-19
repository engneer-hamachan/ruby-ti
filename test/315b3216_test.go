package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test315b3216(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./315b3216.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./315b3216.rb:::13:::Integer"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
