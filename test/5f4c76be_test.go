package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test5f4c76be(t *testing.T) {
	cmd := exec.Command("../ti", "./5f4c76be.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./5f4c76be.rb::31::Union<Integer String>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
