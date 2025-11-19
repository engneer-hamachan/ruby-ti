package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test4e78aa46(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./4e78aa46.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./4e78aa46.rb:::18:::Union<Integer String Array<untyped>>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
