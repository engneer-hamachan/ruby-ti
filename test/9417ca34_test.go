package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test9417ca34(t *testing.T) {
	cmd := exec.Command("../ti", "./9417ca34.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./9417ca34.rb::17::method 'test2' is not defined for Hoge\n./9417ca34.rb::17::method '+' is not defined for Hoge"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
