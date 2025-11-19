package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test90908126(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./90908126.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./90908126.rb:::3:::Integer`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
