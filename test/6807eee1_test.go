package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test6807eee1(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./6807eee1.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./6807eee1.rb:::21:::type mismatch: expected Union<Integer Float>, but got String for Integer.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
