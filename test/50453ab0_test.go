package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test50453ab0(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./50453ab0.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./50453ab0.rb:::13:::type mismatch: expected String, but got Integer for String.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
