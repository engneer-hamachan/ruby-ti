package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test5ae7e1ba(t *testing.T) {
	cmd := exec.Command("../ti", "./5ae7e1ba.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := ""

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
