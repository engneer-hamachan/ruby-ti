package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test1b5a2182(t *testing.T) {
	cmd := exec.Command("../ti", "./1b5a2182.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./1b5a2182.rb::3::Union<Integer String>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
