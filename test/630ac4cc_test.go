package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test630ac4cc(t *testing.T) {
	cmd := exec.Command("../ti", "./630ac4cc.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./630ac4cc.rb::12::type mismatch: expected String, but got Integer for String.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
