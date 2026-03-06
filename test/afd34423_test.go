package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestAfd34423(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./afd34423.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./afd34423.rb:::4:::Bool`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
