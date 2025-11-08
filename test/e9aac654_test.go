package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestE9aac654(t *testing.T) {
	cmd := exec.Command("../ti", "./e9aac654.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./e9aac654.rb:::6:::Unknown
./e9aac654.rb:::18:::Unknown`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
