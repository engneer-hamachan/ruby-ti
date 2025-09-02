package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test0a624557(t *testing.T) {
	cmd := exec.Command("../ti", "./0a624557.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./0a624557.rb::9::Union<Nil Integer>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
