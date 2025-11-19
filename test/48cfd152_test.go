package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test48cfd152(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./48cfd152.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./48cfd152.rb:::16:::class method 'test2' is not defined for MatzSanArigato"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
