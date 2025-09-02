package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test691e1f27(t *testing.T) {
	cmd := exec.Command("../ti", "./691e1f27.rb")


	output, _ := cmd.CombinedOutput()


	expectedOutput := "./691e1f27.rb::20::type mismatch: expected Union<Integer Float>, but got String for Integer.+"


	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
