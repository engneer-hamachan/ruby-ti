package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test8f9129c3(t *testing.T) {
	cmd := exec.Command("../ti", "./8f9129c3.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./8f9129c3.rb:::4:::instance method 'abs' is not defined for String`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
