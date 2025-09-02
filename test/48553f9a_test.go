package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test48553f9a(t *testing.T) {
	cmd := exec.Command("../ti", "./48553f9a.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./48553f9a.rb::4::type mismatch: expected Union<Integer Float>, but got Nil for Integer.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
