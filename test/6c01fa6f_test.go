package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test6c01fa6f(t *testing.T) {
	cmd := exec.Command("../ti", "./6c01fa6f.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./6c01fa6f.rb::6::String
./6c01fa6f.rb::8::Union<Integer Nil>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
