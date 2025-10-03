package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestF37a20d4(t *testing.T) {
	cmd := exec.Command("../ti", "./f37a20d4.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./f37a20d4.rb:::17:::Union<String Integer>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
