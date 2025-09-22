package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test32bad572(t *testing.T) {
	cmd := exec.Command("../ti", "./32bad572.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./32bad572.rb::16::String
./32bad572.rb::17::Union<Nil untyped String>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
