package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestA16ff0ae(t *testing.T) {
	cmd := exec.Command("../ti", "./a16ff0ae.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./a16ff0ae.rb:::21:::Integer`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
