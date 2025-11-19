package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestC9a02fbf(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./c9a02fbf.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./c9a02fbf.rb:::5:::Union<Nil untyped>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
