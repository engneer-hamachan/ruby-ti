package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test4a951f47(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./4a951f47.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./4a951f47.rb:::31:::too few arguments for Piyo.test expected (Integer, unknown)"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
