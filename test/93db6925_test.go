package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test93db6925(t *testing.T) {
	cmd := exec.Command("../ti", "./93db6925.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./93db6925.rb::1::String`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
