package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestAdeeacb3(t *testing.T) {
	cmd := exec.Command("../ti", "./adeeacb3.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./adeeacb3.rb:::128:::method 'read_object' is not defined for JsonSan"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
