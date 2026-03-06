package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test57f232b6(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./57f232b6.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./57f232b6.rb:::6:::Array<untyped>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
