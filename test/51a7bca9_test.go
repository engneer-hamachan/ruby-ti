package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test51a7bca9(t *testing.T) {
	cmd := exec.Command("../ti", "./51a7bca9.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./51a7bca9.rb:::2:::type mismatch: expected String, but got Integer for String.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
