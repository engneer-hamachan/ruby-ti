package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestEc5b367f(t *testing.T) {
	cmd := exec.Command("../ti", "./ec5b367f.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./ec5b367f.rb::2::type mismatch: expected String, but got Integer for String.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
