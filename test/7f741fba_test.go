package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test7f741fba(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./7f741fba.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./7f741fba.rb:::1:::Integer"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
