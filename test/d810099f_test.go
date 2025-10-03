package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestD810099f(t *testing.T) {
	cmd := exec.Command("../ti", "./d810099f.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./d810099f.rb:::13:::type mismatch: expected String, but got Integer for test"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
