package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test61bd7238(t *testing.T) {
	cmd := exec.Command("../ti", "./61bd7238.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./61bd7238.rb::106::Union<Nil Integer Array<String> String Hash>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
