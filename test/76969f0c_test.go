package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test76969f0c(t *testing.T) {
	cmd := exec.Command("../ti", "./76969f0c.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./76969f0c.rb::6::type mismatch: expected String, but got Integer for String.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
