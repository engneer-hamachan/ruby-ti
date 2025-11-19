package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test9bcb8095(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./9bcb8095.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./9bcb8095.rb:::1:::Array<Array<String> Array<Array<untyped> String>>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
