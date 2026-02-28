package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestDaf3de65(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./daf3de65.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./daf3de65.rb:::3:::Integer`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
