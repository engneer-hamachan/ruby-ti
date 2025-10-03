package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestF4e4f41b(t *testing.T) {
	cmd := exec.Command("../ti", "./f4e4f41b.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./f4e4f41b.rb:::9:::method '+' is not defined for Bool"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
