package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test1c424b27(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./1c424b27.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./1c424b27.rb:::3:::Union<Integer String>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
