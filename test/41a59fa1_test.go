package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test41a59fa1(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./41a59fa1.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./41a59fa1.rb:::2:::type mismatch: expected Union<Integer Float>, but got NilClass for Integer.+\n./41a59fa1.rb:::4:::type mismatch: expected Union<Integer Float>, but got String for Integer.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
