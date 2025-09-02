package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test3ffcef84(t *testing.T) {
	cmd := exec.Command("../ti", "./3ffcef84.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./3ffcef84.rb::11::Hoge.a is read only"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
