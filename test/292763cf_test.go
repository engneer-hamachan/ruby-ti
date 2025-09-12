package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test292763cf(t *testing.T) {
	cmd := exec.Command("../ti", "./292763cf.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./292763cf.rb::3::Hash`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
