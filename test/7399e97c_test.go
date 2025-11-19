package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test7399e97c(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./7399e97c.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./7399e97c.rb:::3:::String"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
