package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestD3e75739(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./d3e75739.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./d3e75739.rb:::16:::String"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
