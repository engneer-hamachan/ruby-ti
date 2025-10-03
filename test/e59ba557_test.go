package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestE59ba557(t *testing.T) {
	cmd := exec.Command("../ti", "./e59ba557.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./e59ba557.rb:::9:::Array<String>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
