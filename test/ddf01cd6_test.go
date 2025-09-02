package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestDdf01cd6(t *testing.T) {
	cmd := exec.Command("../ti", "./ddf01cd6.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./ddf01cd6.rb::2::Union<Nil Integer>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
