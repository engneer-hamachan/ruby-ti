package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestFf73942f(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./ff73942f.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./ff73942f.rb:::12:::String"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
