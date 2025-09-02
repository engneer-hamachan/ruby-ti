package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test151387dd(t *testing.T) {
	cmd := exec.Command("../ti", "./151387dd.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./151387dd.rb::9::type mismatch: expected String, but got Integer for String.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
