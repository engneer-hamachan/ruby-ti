package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test9c80e61f(t *testing.T) {
	cmd := exec.Command("../ti", "./9c80e61f.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./9c80e61f.rb::18::type mismatch: expected Union<Integer Float>, but got String for Integer.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
