package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test388a8217(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./388a8217.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./388a8217.rb:::3:::Array<Integer>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
