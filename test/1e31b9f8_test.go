package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test1e31b9f8(t *testing.T) {
	cmd := exec.Command("../ti", "./1e31b9f8.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./1e31b9f8.rb::8::Union<Integer Nil>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
