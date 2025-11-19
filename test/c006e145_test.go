package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestC006e145(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./c006e145.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./c006e145.rb:::5:::Union<Integer String>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
