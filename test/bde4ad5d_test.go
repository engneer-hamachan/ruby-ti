package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestBde4ad5d(t *testing.T) {
	cmd := exec.Command("../ti", "./bde4ad5d.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./bde4ad5d.rb::11::method 'test' is not defined for Identifier
./bde4ad5d.rb::11::Unknown`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
