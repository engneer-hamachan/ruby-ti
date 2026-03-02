package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test027022eb(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./027022eb.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./027022eb.rb:::13:::Integer`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
