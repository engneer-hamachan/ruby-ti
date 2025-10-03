package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestEbb4ed62(t *testing.T) {
	cmd := exec.Command("../ti", "./ebb4ed62.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./ebb4ed62.rb:::12:::too few arguments for Hoge.test2 expected (Integer)"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
