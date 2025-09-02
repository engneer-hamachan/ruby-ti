package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test0b53545a(t *testing.T) {
	cmd := exec.Command("../ti", "./0b53545a.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./0b53545a.rb::2::method 'abs' is not defined for String\n./0b53545a.rb::7::type mismatch: expected String, but got Integer for String.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
