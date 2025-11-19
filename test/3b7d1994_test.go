package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test3b7d1994(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./3b7d1994.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./3b7d1994.rb:::6:::Union<Integer String>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
