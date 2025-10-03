package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test03d9258d(t *testing.T) {
	cmd := exec.Command("../ti", "./03d9258d.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./03d9258d.rb:::1:::Array<Integer>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
