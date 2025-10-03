package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test0aec9d16(t *testing.T) {
	cmd := exec.Command("../ti", "./0aec9d16.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./0aec9d16.rb:::12:::Array<String>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
