package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestDbca74e4(t *testing.T) {
	cmd := exec.Command("../ti", "./dbca74e4.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./dbca74e4.rb::2::syntax errror. test is define multiple '*'"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
