package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test92f49cdf(t *testing.T) {
	cmd := exec.Command("../ti", "./92f49cdf.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./92f49cdf.rb:::6:::Range"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
