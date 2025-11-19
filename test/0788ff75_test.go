package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test0788ff75(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./0788ff75.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./0788ff75.rb:::2:::Union<Integer String>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
