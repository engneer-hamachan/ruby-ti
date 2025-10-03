package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestB93673b8(t *testing.T) {
	cmd := exec.Command("../ti", "./b93673b8.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./b93673b8.rb:::10:::type mismatch: expected String, but got Integer for String.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
