package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestAd4a9fbf(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./ad4a9fbf.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./ad4a9fbf.rb:::27:::String`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
