package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test400d0c02(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./400d0c02.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./400d0c02.rb:::2:::Float`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
