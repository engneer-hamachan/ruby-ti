package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestCb292fc2(t *testing.T) {
	cmd := exec.Command("../ti", "./cb292fc2.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./cb292fc2.rb:::2:::Integer"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
