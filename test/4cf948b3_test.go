package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test4cf948b3(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./4cf948b3.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./4cf948b3.rb:::10:::Integer
./4cf948b3.rb:::11:::String`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
