package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestAd0abdd9(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./ad0abdd9.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./ad0abdd9.rb:::5:::too few arguments for test expected (unknown)`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
