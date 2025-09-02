package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test398fa758(t *testing.T) {
	cmd := exec.Command("../ti", "./398fa758.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./398fa758.rb::2::Nil\n./398fa758.rb::10::type mismatch: expected Union<Integer Float>, but got String for Integer.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
