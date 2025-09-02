package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test02dfa6c0(t *testing.T) {
	cmd := exec.Command("../ti", "./02dfa6c0.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./02dfa6c0.rb::3::too many arguments for String.to_i"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
