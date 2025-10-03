package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test3de5c8d2(t *testing.T) {
	cmd := exec.Command("../ti", "./3de5c8d2.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./3de5c8d2.rb::8::Hoge inner Fuga`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
