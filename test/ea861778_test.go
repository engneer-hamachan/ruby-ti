package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestEa861778(t *testing.T) {
	cmd := exec.Command("../ti", "./ea861778.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./ea861778.rb::1::Union<Integer Nil>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
