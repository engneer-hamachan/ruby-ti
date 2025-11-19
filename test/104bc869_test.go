package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test104bc869(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./104bc869.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./104bc869.rb:::11:::ActiveRecord::Relation`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
