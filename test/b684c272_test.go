package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestB684c272(t *testing.T) {
	cmd := exec.Command("../ti", "./b684c272.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./b684c272.rb::17::Union<String Nil>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
