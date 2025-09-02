package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test0ee9a5ff(t *testing.T) {
	cmd := exec.Command("../ti", "./0ee9a5ff.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./0ee9a5ff.rb::11::method 'abs' is not defined for Identifier"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
