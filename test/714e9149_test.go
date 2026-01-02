package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test714e9149(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./714e9149.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./714e9149.rb:::2:::Union<NilClass Integer>\n./714e9149.rb:::10:::type mismatch: expected Union<Integer Float>, but got String for Integer.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
