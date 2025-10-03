package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test8c22a83f(t *testing.T) {
	cmd := exec.Command("../ti", "./8c22a83f.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./8c22a83f.rb:::1:::Array<Array<untyped>>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
