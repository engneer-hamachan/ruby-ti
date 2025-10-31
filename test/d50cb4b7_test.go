package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestD50cb4b7(t *testing.T) {
	cmd := exec.Command("../ti", "./d50cb4b7.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./d50cb4b7.rb:::13:::instance method 'fuga' is not defined for Hoge"
	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
