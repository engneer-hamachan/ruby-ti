package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test7092cc4a(t *testing.T) {
	cmd := exec.Command("../ti", "./7092cc4a.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := ""

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
