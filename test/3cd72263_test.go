package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test3cd72263(t *testing.T) {
	cmd := exec.Command("../ti", "./3cd72263.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./3cd72263.rb:::2:::type mismatch: expected Union<Integer Float>, but got Union<Integer String> for Integer.+`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
