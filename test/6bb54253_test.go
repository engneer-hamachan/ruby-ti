package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test6bb54253(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./6bb54253.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./6bb54253.rb:::14:::instance method 'piyo' is not defined for Hoge"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
