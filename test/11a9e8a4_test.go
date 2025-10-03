package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test11a9e8a4(t *testing.T) {
	cmd := exec.Command("../ti", "./11a9e8a4.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./11a9e8a4.rb:::7:::Integer"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
