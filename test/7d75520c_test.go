package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test7d75520c(t *testing.T) {
	cmd := exec.Command("../ti", "./7d75520c.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./7d75520c.rb::5::Array<Integer>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
