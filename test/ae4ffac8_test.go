package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestAe4ffac8(t *testing.T) {
	cmd := exec.Command("../ti", "./ae4ffac8.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./ae4ffac8.rb:::8:::Integer`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
