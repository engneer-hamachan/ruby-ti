package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test6875f4c2(t *testing.T) {
	cmd := exec.Command("../ti", "./6875f4c2.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./6875f4c2.rb::11::method 'test' is not defined for Hoge"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
