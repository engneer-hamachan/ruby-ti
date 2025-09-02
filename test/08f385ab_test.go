package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test08f385ab(t *testing.T) {
	cmd := exec.Command("../ti", "./08f385ab.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./08f385ab.rb::9::Union<Integer Nil>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
