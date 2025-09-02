package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test9935e210(t *testing.T) {
	cmd := exec.Command("../ti", "./9935e210.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./9935e210.rb::5::type mismatch: expected Integer, but got String for hoge"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
