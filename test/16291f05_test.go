package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test16291f05(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./16291f05.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./16291f05.rb:::9:::untyped`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
