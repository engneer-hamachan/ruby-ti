package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestB67dc016(t *testing.T) {
	cmd := exec.Command("../ti", "./b67dc016.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./b67dc016.rb::1::Float"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
