package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test77b9207b(t *testing.T) {
	cmd := exec.Command("../ti", "./77b9207b.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./77b9207b.rb:::21:::String"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
