package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestFe07c496(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./fe07c496.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./fe07c496.rb:::5:::String`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
