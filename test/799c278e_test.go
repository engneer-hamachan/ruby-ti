package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test799c278e(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./799c278e.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./799c278e.rb:::7:::type mismatch: expected Union<Integer Float>, but got Union<Integer String> for Integer.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
