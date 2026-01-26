package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test94b434b4(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./94b434b4.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./94b434b4.rb:::8:::Union<String NilClass>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
