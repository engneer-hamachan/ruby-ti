package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test8cb3be8c(t *testing.T) {
	cmd := exec.Command("../ti", "./8cb3be8c.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./8cb3be8c.rb:::4:::type mismatch: expected Array<untyped>, but got Integer for Array.concat\n./8cb3be8c.rb:::6:::type mismatch: expected Union<Integer Float>, but got String for Integer.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
