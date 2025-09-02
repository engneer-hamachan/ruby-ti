package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test6358c3f1(t *testing.T) {
	cmd := exec.Command("../ti", "./6358c3f1.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./6358c3f1.rb::7::type mismatch: expected Block, but got Integer for Hash.each"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
