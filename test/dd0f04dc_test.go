package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestDd0f04dc(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./dd0f04dc.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./dd0f04dc.rb:::3:::Bot"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
