package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestAfe1372d(t *testing.T) {
	cmd := exec.Command("../ti", "./afe1372d.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./afe1372d.rb::13::Array<String Integer>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
