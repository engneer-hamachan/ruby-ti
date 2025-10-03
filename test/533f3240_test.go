package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test533f3240(t *testing.T) {
	cmd := exec.Command("../ti", "./533f3240.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./533f3240.rb:::5:::Integer"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
