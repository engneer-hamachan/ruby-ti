package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestDde150c2(t *testing.T) {
	cmd := exec.Command("../ti", "./dde150c2.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./dde150c2.rb:::5:::type mismatch: expected String, but got Integer for String.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
