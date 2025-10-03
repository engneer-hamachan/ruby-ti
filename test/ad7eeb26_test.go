package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestAd7eeb26(t *testing.T) {
	cmd := exec.Command("../ti", "./ad7eeb26.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./ad7eeb26.rb:::5:::type mismatch: expected String, but got Integer for String.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
