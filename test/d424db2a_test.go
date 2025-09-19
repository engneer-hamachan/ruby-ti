package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestD424db2a(t *testing.T) {
	cmd := exec.Command("../ti", "./d424db2a.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./d424db2a.rb::12::Union<Nil Integer>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
