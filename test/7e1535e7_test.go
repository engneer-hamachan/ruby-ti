package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test7e1535e7(t *testing.T) {
	cmd := exec.Command("../ti", "./7e1535e7.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := ``

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
