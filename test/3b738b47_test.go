package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test3b738b47(t *testing.T) {
	cmd := exec.Command("../ti", "./3b738b47.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./3b738b47.rb:::2:::var77 is not defined expected (Union<Integer Float>)\n./3b738b47.rb:::2:::type mismatch: expected Union<Integer Float>, but got Keyword for Integer.+\n./3b738b47.rb:::4:::type mismatch: expected Union<Integer Float>, but got Nil for Integer.+\n./3b738b47.rb:::7:::Integer is extra argument"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
