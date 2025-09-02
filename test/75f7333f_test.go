package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test75f7333f(t *testing.T) {
	cmd := exec.Command("../ti", "./75f7333f.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./75f7333f.rb::8::method '+' is not defined for Identifier\n./75f7333f.rb::11::too few arguments for test2 expected (?)"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
