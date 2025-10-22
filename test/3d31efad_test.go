package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test3d31efad(t *testing.T) {
	cmd := exec.Command("../ti", "./3d31efad.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./3d31efad.rb:::5:::String
./3d31efad.rb:::6:::Integer`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
