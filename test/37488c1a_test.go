package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test37488c1a(t *testing.T) {
	cmd := exec.Command("../ti", "./37488c1a.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./37488c1a.rb::3::untyped
./37488c1a.rb::7::untyped`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
