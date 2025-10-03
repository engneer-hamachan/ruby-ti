package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test83a7cf9d(t *testing.T) {
	cmd := exec.Command("../ti", "./83a7cf9d.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./83a7cf9d.rb:::2:::type mismatch: expected Union<Integer Float>, but got String for Integer.+\n./83a7cf9d.rb:::2:::String"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
