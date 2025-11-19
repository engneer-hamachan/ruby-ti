package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test2197a40b(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./2197a40b.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./2197a40b.rb:::3:::Union<Integer Nil>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
