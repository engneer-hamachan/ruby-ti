package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestEfb7f957(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./efb7f957.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./efb7f957.rb:::2:::Integer
./efb7f957.rb:::3:::Hash`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
