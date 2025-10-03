package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test43476ab6(t *testing.T) {
	cmd := exec.Command("../ti", "./43476ab6.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./43476ab6.rb:::9:::Array<Integer>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
