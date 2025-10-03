package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestE75012f4(t *testing.T) {
	cmd := exec.Command("../ti", "./e75012f4.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./e75012f4.rb:::21:::type mismatch: expected Union<Integer Float>, but got String for Integer.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
