package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test5f7ebd48(t *testing.T) {
	cmd := exec.Command("../ti", "./5f7ebd48.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./5f7ebd48.rb:::19:::instance method 'hoge' is not defined for Hoge"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
