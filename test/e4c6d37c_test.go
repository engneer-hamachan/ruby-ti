package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestE4c6d37c(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./e4c6d37c.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./e4c6d37c.rb:::1:::Union<Integer String>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
