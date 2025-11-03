package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test65b3b849(t *testing.T) {
	cmd := exec.Command("../ti", "./65b3b849.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./65b3b849.rb:::6:::too few arguments for Hoge.write expected (Unknown)`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
