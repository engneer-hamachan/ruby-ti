package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestF23ee39b(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./f23ee39b.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./f23ee39b.rb:::4:::type mismatch: expected Union<Integer Float>, but got Nil for Integer.+\n./f23ee39b.rb:::7:::q is not defined expected (Integer, unknown, a: Integer, b: Integer, c: optional Nil)"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
