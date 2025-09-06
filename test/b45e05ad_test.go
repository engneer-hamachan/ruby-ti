package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestB45e05ad(t *testing.T) {
	cmd := exec.Command("../ti", "./b45e05ad.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./b45e05ad.rb::3::String`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
