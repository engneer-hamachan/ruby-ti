package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test25939376(t *testing.T) {
	cmd := exec.Command("../ti", "./25939376.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./25939376.rb:::13:::instance method '+' is not defined for Hash"
	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
