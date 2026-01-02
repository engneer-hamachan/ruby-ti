package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestAa8cbd99(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./aa8cbd99.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./aa8cbd99.rb:::1:::Union<Integer NilClass String>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
