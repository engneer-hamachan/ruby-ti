package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestAd95e7e1(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./ad95e7e1.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := ``

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
