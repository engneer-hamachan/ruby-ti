package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test804f9c85(t *testing.T) {
	cmd := exec.Command("../ti", "./804f9c85.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./804f9c85.rb:::2:::type mismatch: expected Union<Integer Float>, but got String for Integer.+\n./804f9c85.rb:::5:::method 'any' is not defined for Integer"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
