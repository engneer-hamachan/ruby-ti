package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test39429dff(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./39429dff.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./39429dff.rb:::8:::String"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
