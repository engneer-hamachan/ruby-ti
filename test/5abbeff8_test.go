package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test5abbeff8(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./5abbeff8.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./5abbeff8.rb:::27:::type mismatch: expected Array<untyped>, but got Integer for Array.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
