package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test9cc57b5a(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./9cc57b5a.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./9cc57b5a.rb:::16:::String"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
