package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestB191348e(t *testing.T) {
	cmd := exec.Command("../ti", "./b191348e.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./b191348e.rb:::8:::Array<Array<Symbol Integer String>>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
