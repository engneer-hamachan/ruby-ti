package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestB1393fe9(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./b1393fe9.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./b1393fe9.rb:::6:::Integer"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
