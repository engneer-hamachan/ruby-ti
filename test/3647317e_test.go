package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test3647317e(t *testing.T) {
	cmd := exec.Command("../ti", "./3647317e.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./3647317e.rb:::14:::Hash"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
