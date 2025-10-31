package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestDa6bb92e(t *testing.T) {
	cmd := exec.Command("../ti", "./da6bb92e.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./da6bb92e.rb:::13:::instance method '+' is not defined for Hoge"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
