package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestE262461c(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./e262461c.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./e262461c.rb:::6:::instance method 'uniq' is not defined for NilClass"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
