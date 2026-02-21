package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test6a01e017(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./6a01e017.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./6a01e017.rb:::8:::String"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
