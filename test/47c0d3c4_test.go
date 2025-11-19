package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test47c0d3c4(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./47c0d3c4.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./47c0d3c4.rb:::2:::Union<Integer String>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
