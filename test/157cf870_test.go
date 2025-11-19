package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test157cf870(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./157cf870.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./157cf870.rb:::5:::type mismatch: expected Union<Integer Float>, but got String for Integer.+"
	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
