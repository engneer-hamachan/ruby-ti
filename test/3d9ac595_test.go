package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test3d9ac595(t *testing.T) {
	cmd := exec.Command("../ti", "./3d9ac595.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./3d9ac595.rb:::3:::String"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
