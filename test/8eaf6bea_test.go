package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test8eaf6bea(t *testing.T) {
	cmd := exec.Command("../ti", "./8eaf6bea.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./8eaf6bea.rb::4::Integer
./8eaf6bea.rb::5::Array<untyped>
./8eaf6bea.rb::6::String
./8eaf6bea.rb::7::String
./8eaf6bea.rb::8::Integer`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
