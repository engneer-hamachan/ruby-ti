package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test1bb99c0a(t *testing.T) {
	cmd := exec.Command("../ti", "./1bb99c0a.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./1bb99c0a.rb::3::Point
./1bb99c0a.rb::12::Point`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
