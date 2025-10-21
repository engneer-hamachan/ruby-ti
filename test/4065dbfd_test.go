package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test4065dbfd(t *testing.T) {
	cmd := exec.Command("../ti", "./4065dbfd.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./4065dbfd.rb:::16:::String`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
