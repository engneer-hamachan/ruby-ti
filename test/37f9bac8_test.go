package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test37f9bac8(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./37f9bac8.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./37f9bac8.rb:::12:::Array<Integer>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
