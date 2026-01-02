package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestFe7eb355(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./fe7eb355.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./fe7eb355.rb:::7:::Union<Integer String NilClass>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
