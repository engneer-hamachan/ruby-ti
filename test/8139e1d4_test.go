package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test8139e1d4(t *testing.T) {
	cmd := exec.Command("../ti", "./8139e1d4.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./8139e1d4.rb::3::Array<Integer>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
