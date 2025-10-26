package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestC56c6b5b(t *testing.T) {
	cmd := exec.Command("../ti", "./c56c6b5b.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./c56c6b5b.rb:::2:::Union<Integer Float>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
