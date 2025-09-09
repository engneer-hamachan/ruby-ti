package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test93f6dc62(t *testing.T) {
	cmd := exec.Command("../ti", "./93f6dc62.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./93f6dc62.rb::2::method 'size' is not defined for Identifier
./93f6dc62.rb::10::untyped`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
