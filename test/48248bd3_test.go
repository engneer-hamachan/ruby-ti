package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test48248bd3(t *testing.T) {
	cmd := exec.Command("../ti", "./48248bd3.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./48248bd3.rb::3::type mismatch: expected Array<untyped>, but got Integer for Array.+"
	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
