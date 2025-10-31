package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test80339eb6(t *testing.T) {
	cmd := exec.Command("../ti", "./80339eb6.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./80339eb6.rb:::2:::instance method 'hoge' is not defined for Hoge`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
