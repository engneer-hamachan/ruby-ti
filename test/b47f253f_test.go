package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestB47f253f(t *testing.T) {
	cmd := exec.Command("../ti", "./b47f253f.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./b47f253f.rb::9::type mismatch: expected String, but got Integer for String.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
