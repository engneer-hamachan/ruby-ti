package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test13f156e8(t *testing.T) {
	cmd := exec.Command("../ti", "./13f156e8.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./13f156e8.rb::1::Range"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
