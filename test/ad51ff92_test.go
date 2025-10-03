package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestAd51ff92(t *testing.T) {
	cmd := exec.Command("../ti", "./ad51ff92.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./ad51ff92.rb:::3:::type mismatch: expected Union<Integer Float>, but got String for Integer.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
