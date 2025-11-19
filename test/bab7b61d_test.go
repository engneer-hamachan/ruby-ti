package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestBab7b61d(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./bab7b61d.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./bab7b61d.rb:::1:::Union<Integer String>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
