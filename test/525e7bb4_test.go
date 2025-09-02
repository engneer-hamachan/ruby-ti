package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test525e7bb4(t *testing.T) {
	cmd := exec.Command("../ti", "./525e7bb4.rb")


	output, _ := cmd.CombinedOutput()


	expectedOutput := "./525e7bb4.rb::5::type mismatch: expected Union<Integer Float>, but got String for Integer.+"


	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
