package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test9fb4ff89(t *testing.T) {
	cmd := exec.Command("../ti", "./9fb4ff89.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./9fb4ff89.rb:::5:::Union<Nil String>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
