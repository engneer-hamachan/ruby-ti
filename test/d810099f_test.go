package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestD810099f(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./d810099f.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./d810099f.rb:::8:::type mismatch: expected String, but got Integer for String.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
