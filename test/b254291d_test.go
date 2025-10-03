package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestB254291d(t *testing.T) {
	cmd := exec.Command("../ti", "./b254291d.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./b254291d.rb:::2:::Hash"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
