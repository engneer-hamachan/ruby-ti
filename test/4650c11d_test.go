package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test4650c11d(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./4650c11d.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./4650c11d.rb:::24:::String
./4650c11d.rb:::25:::Array<untyped>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
