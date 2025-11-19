package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test1ee4317b(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./1ee4317b.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./1ee4317b.rb:::2:::type mismatch: expected Union<Integer Float>, but got Nil for Integer.+\n./1ee4317b.rb:::4:::type mismatch: expected String, but got Integer for String.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
