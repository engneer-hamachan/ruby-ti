package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test6d715d7b(t *testing.T) {
	cmd := exec.Command("../ti", "./6d715d7b.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./6d715d7b.rb::1::Array<Integer>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
