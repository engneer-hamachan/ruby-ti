package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test6d192e15(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./6d192e15.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./6d192e15.rb:::2:::Bool
./6d192e15.rb:::3:::Integer`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
