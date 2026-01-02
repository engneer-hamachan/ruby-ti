package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test81c8bd31(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./81c8bd31.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./81c8bd31.rb:::3:::Union<Array<String> NilClass>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
