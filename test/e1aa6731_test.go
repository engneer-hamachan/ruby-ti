package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestE1aa6731(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./e1aa6731.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./e1aa6731.rb:::4:::type missmatch error: given String expected Class for is_a?\n./e1aa6731.rb:::5:::instance method '+' is not defined for Nil\n./e1aa6731.rb:::9:::Unknown"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
