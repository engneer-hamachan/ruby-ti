package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test53b5eed3(t *testing.T) {
	cmd := exec.Command("../ti", "./53b5eed3.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./53b5eed3.rb:::14:::String"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
