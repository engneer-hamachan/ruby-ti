package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test75f7333f(t *testing.T) {
	cmd := exec.Command("../ti", "./75f7333f.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./75f7333f.rb:::11:::too few arguments for test2 expected (unknown)"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
