package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test9efed07b(t *testing.T) {
	cmd := exec.Command("../ti", "./9efed07b.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./9efed07b.rb:::13:::String"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
