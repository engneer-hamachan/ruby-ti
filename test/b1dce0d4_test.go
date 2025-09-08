package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestB1dce0d4(t *testing.T) {
	cmd := exec.Command("../ti", "./b1dce0d4.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./b1dce0d4.rb::7::Array<Integer String>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
