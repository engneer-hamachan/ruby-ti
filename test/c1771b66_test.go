package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestC1771b66(t *testing.T) {
	cmd := exec.Command("../ti", "./c1771b66.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./c1771b66.rb:::18:::instance method 'test' is not defined for Hoge"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
