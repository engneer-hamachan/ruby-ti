package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test0e00c5d3(t *testing.T) {
	cmd := exec.Command("../ti", "./0e00c5d3.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./0e00c5d3.rb::6::Union<Integer Nil>
./0e00c5d3.rb::8::String`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
