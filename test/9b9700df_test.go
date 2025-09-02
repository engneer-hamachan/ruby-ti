package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test9b9700df(t *testing.T) {
	cmd := exec.Command("../ti", "./9b9700df.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./9b9700df.rb::28::method 'test2' is not defined for Fuga\n./9b9700df.rb::28::method '+' is not defined for Fuga"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
