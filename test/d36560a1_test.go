package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestD36560a1(t *testing.T) {
	cmd := exec.Command("../ti", "./d36560a1.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./d36560a1.rb::15::method 'new' is not defined for Fuga\n./d36560a1.rb::17::String"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
