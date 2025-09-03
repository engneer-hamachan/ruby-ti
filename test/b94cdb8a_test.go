package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestB94cdb8a(t *testing.T) {
	cmd := exec.Command("../ti", "./b94cdb8a.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./b94cdb8a.rb::15::Union<String Integer>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
