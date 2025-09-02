package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test19fec8c9(t *testing.T) {
	cmd := exec.Command("../ti", "./19fec8c9.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./19fec8c9.rb::9::type mismatch: expected String, but got Integer for String.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
