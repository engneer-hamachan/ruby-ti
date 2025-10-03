package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestC55aed5d(t *testing.T) {
	cmd := exec.Command("../ti", "./c55aed5d.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./c55aed5d.rb:::14:::Nil"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
