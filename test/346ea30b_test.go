package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test346ea30b(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./346ea30b.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./346ea30b.rb:::6:::String"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
