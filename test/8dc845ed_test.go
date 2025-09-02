package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test8dc845ed(t *testing.T) {
	cmd := exec.Command("../ti", "./8dc845ed.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./8dc845ed.rb::11::String"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
