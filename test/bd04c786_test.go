package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestBd04c786(t *testing.T) {
	cmd := exec.Command("../ti", "./bd04c786.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./bd04c786.rb::17::method 'test2' is not defined for Hoge
./bd04c786.rb::17::Hoge`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
