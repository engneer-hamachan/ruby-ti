package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test8a9e3957(t *testing.T) {
	cmd := exec.Command("../ti", "./8a9e3957.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./8a9e3957.rb::3::type mismatch: expected String, but got Integer for String.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
