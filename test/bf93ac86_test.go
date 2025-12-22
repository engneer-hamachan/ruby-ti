package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestBf93ac86(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./bf93ac86.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./bf93ac86.rb:::3:::2 is not Integer\n./bf93ac86.rb:::3:::String"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
