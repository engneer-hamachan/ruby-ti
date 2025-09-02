package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test9d472ec3(t *testing.T) {
	cmd := exec.Command("../ti", "./9d472ec3.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./9d472ec3.rb::10::type mismatch: expected String, but got Integer for String.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
