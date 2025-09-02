package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test0f21dcd5(t *testing.T) {
	cmd := exec.Command("../ti", "./0f21dcd5.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./0f21dcd5.rb::1::type mismatch: expected Union<Integer Float>, but got Bool for Integer.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
