package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test530fcf92(t *testing.T) {
	cmd := exec.Command("../ti", "./530fcf92.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./530fcf92.rb::4::type mismatch: expected Array<untyped>, but got Integer for Array.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
