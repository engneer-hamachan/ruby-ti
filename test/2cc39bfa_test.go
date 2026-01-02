package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test2cc39bfa(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./2cc39bfa.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./2cc39bfa.rb:::13:::Union<Integer NilClass>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
