package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test8d999578(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./8d999578.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./8d999578.rb:::9:::instance method '+' is not defined for Bool`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
