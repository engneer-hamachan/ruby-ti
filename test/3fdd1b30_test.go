package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test3fdd1b30(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./3fdd1b30.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./3fdd1b30.rb:::6:::instance method 'chars' is not defined for Nil"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
