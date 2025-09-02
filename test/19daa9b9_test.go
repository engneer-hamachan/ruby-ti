package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test19daa9b9(t *testing.T) {
	cmd := exec.Command("../ti", "./19daa9b9.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./19daa9b9.rb::7::Array<String Symbol>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
