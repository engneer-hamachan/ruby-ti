package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestA88aaf59(t *testing.T) {
	cmd := exec.Command("../ti", "./a88aaf59.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./a88aaf59.rb::3::Union<Nil String>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
