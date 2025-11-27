package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test23ee4148(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./23ee4148.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./23ee4148.rb:::22:::Integer
./23ee4148.rb:::23:::String`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
