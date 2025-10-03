package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestF9778acf(t *testing.T) {
	cmd := exec.Command("../ti", "./f9778acf.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./f9778acf.rb:::13:::type mismatch: expected String, but got Integer for String.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
