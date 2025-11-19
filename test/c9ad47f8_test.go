package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestC9ad47f8(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./c9ad47f8.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./c9ad47f8.rb:::27:::String`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
