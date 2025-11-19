package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestC2bb6da8(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./c2bb6da8.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./c2bb6da8.rb:::35:::String"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
