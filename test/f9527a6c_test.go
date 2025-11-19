package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestF9527a6c(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./f9527a6c.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./f9527a6c.rb:::4:::String
./f9527a6c.rb:::9:::Integer`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
