package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestD0abc991(t *testing.T) {
	cmd := exec.Command("../ti", "./d0abc991.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./d0abc991.rb:::4:::Integer"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
