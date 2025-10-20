package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test5ed8a2d1(t *testing.T) {
	cmd := exec.Command("../ti", "./5ed8a2d1.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./5ed8a2d1.rb:::22:::Piyo`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
