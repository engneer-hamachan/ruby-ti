package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test11617386(t *testing.T) {
	cmd := exec.Command("../ti", "./11617386.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./11617386.rb:::5:::Union<Nil untyped>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
