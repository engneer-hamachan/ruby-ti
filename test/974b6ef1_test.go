package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test974b6ef1(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./974b6ef1.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./974b6ef1.rb:::13:::type mismatch: expected String, but got Integer for String.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
