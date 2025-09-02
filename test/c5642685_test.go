package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestC5642685(t *testing.T) {
	cmd := exec.Command("../ti", "./c5642685.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./c5642685.rb::16::Union<Nil Integer Array<untyped> String>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
