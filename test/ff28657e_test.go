package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestFf28657e(t *testing.T) {
	cmd := exec.Command("../ti", "./ff28657e.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./ff28657e.rb::14::Union<Integer String>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
