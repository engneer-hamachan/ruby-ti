package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test57b97647(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./57b97647.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./57b97647.rb:::1:::Integer"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
