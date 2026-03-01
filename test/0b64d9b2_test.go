package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test0b64d9b2(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./0b64d9b2.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./0b64d9b2.rb:::7:::Integer`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
