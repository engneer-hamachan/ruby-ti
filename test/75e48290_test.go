package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test75e48290(t *testing.T) {
	cmd := exec.Command("../ti", "./75e48290.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./75e48290.rb:::2:::Integer
./75e48290.rb:::3:::Unknown
./75e48290.rb:::6:::expected keyvalue argument for **kwargs parameter`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
