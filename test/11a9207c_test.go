package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test11a9207c(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./11a9207c.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./11a9207c.rb:::8:::instance method '+' is not defined for Block"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
