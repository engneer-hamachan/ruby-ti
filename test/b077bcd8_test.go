package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestB077bcd8(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./b077bcd8.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./b077bcd8.rb:::6:::Nil"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
