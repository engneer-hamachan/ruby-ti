package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test4235c46a(t *testing.T) {
	cmd := exec.Command("../ti", "./4235c46a.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./4235c46a.rb:::3:::type mismatch: expected Array<untyped>, but got Union<Integer Hoge> for Array.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
