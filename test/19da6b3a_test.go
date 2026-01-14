package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test19da6b3a(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./19da6b3a.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./19da6b3a.rb:::9:::untyped`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
