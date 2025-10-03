package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test2dad1a74(t *testing.T) {
	cmd := exec.Command("../ti", "./2dad1a74.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./2dad1a74.rb:::21:::Student`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
