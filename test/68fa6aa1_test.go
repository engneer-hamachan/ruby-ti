package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test68fa6aa1(t *testing.T) {
	cmd := exec.Command("../ti", "./68fa6aa1.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./68fa6aa1.rb::8::method '+' is not defined for Nil"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
