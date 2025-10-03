package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test9bb42769(t *testing.T) {
	cmd := exec.Command("../ti", "./9bb42769.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./9bb42769.rb:::3:::method 'to_i' is not defined for Array"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
