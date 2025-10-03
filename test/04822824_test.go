package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test04822824(t *testing.T) {
	cmd := exec.Command("../ti", "./04822824.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./04822824.rb:::4:::too many arguments for test"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
