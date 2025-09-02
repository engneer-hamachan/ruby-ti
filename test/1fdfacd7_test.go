package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test1fdfacd7(t *testing.T) {
	cmd := exec.Command("../ti", "./1fdfacd7.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./1fdfacd7.rb::4::Union<String Nil>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
