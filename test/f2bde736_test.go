package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestF2bde736(t *testing.T) {
	cmd := exec.Command("../ti", "./f2bde736.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./f2bde736.rb:::15:::String`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
