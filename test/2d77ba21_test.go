package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test2d77ba21(t *testing.T) {
	cmd := exec.Command("../ti", "./2d77ba21.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./2d77ba21.rb::9::type mismatch: expected String, but got Integer for String.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
