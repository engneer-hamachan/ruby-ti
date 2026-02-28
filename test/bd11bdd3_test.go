package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestBd11bdd3(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./bd11bdd3.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./bd11bdd3.rb:::4:::Integer`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
