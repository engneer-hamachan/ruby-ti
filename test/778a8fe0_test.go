package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test778a8fe0(t *testing.T) {
	cmd := exec.Command("../ti", "./778a8fe0.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./778a8fe0.rb::15::Array<Integer>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
