package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test839ac2d1(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./839ac2d1.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./839ac2d1.rb:::1:::Array<Array<Integer String> Array<Integer>>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
