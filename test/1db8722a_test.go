package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test1db8722a(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./1db8722a.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./1db8722a.rb:::7:::String"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
