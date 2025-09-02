package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestAd840087(t *testing.T) {
	cmd := exec.Command("../ti", "./ad840087.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./ad840087.rb::14::method 'abs' is not defined for String"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
