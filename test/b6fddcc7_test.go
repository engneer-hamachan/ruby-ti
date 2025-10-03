package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestB6fddcc7(t *testing.T) {
	cmd := exec.Command("../ti", "./b6fddcc7.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./b6fddcc7.rb:::8:::too many arguments for Integer.to_s"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
