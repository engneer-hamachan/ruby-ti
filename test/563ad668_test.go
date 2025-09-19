package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test563ad668(t *testing.T) {
	cmd := exec.Command("../ti", "./563ad668.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./563ad668.rb::130::Union<Nil Integer Array<String> String Hash>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
