package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test9d79e078(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./9d79e078.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./9d79e078.rb:::12:::type mismatch: expected String, but got Integer for String.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
